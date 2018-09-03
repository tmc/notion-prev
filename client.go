package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tmc/notion/notiontypes"
)

const defaultBaseURL = "https://www.notion.so/api/v3/"

// Client is the primary type that implements an interface to the notion.so API.
type Client struct {
	baseURL string
	token   string
	client  *http.Client
	logger  Logger
}

// NewClient initializes a new Client.
func NewClient(opts ...ClientOption) (*Client, error) {
	c := &Client{
		baseURL: defaultBaseURL,
		logger:  &WrapLogrus{logrus.New()},
	}
	for _, o := range opts {
		o(c)
	}
	if c.client == nil {
		c.client = http.DefaultClient
	}
	return c, nil
}

func (c *Client) url(path string) string {
	return fmt.Sprintf("%s%s", c.baseURL, path)
}

func (c *Client) get(pattern string, args ...interface{}) ([]byte, error) {
	return c.do("GET", nil, pattern, args...)
}

func (c *Client) post(payload interface{}, pattern string, args ...interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, err
	}
	c.logger.WithField("fn", "post").Debugln(buf.String())
	return c.do("POST", buf, pattern, args...)
}

func (c *Client) do(method string, body io.Reader, pattern string, args ...interface{}) ([]byte, error) {
	path := c.url(fmt.Sprintf(pattern, args...))
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}
	req.Header.Set("cookie", fmt.Sprintf("token=%v", c.token))
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "performing request")
	}
	defer resp.Body.Close()
	logger := c.logger.WithField("method", method).WithField("path", path).WithField("status_code", resp.StatusCode)
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warnln("error reading body")
		return nil, err
	}
	logger.WithField("body", string(buf)).Debugln("api call finished")
	if resp.StatusCode != http.StatusOK {
		return buf, &Error{
			URL:        path,
			StatusCode: resp.StatusCode,
			Body:       string(buf),
		}
	}
	return buf, nil
}

type loadPageChunkRequest struct {
	PageID          string `json:"pageId"`
	Limit           int64  `json:"limit,omitempty"`
	Cursor          Cursor `json:"cursor"`
	VerticalColumns bool   `json:"verticalColumns"`
}

type loadPageChunkResponse struct {
	RecordMap notiontypes.RecordMap `json:"recordMap"`
	Cursor    Cursor                `json:"cursor"`
}

// GetPage returns a Page given an id.
func (c *Client) GetPage(pageID string) (*Page, error) {
	lp := loadPageChunkRequest{
		PageID: pageID,
		Limit:  50,
		Cursor: Cursor{
			Stack: [][]StackPosition{},
		},
	}
	results := []notiontypes.RecordMap{}
	for {
		r := &loadPageChunkResponse{}
		b, err := c.post(lp, "loadPageChunk")
		if err != nil {
			return nil, err
		}
		c.logger.WithField("pageID", pageID).Debugln(string(b))
		if err := json.Unmarshal(b, r); err != nil {
			return nil, errors.Wrap(err, "unmarshaling loadPageChunkResponse")
		}
		results = append(results, r.RecordMap)
		lp.Cursor = r.Cursor
		if len(r.Cursor.Stack) == 0 {
			break
		}
	}
	return c.parsePageFromRecordMaps(pageID, results)
}

func mergeRecordMaps(rms ...notiontypes.RecordMap) (notiontypes.RecordMap, error) {
	result := notiontypes.RecordMap{
		Blocks:          make(map[string]*notiontypes.BlockWithRole, 50*len(rms)-1),
		Space:           make(map[string]*notiontypes.SpaceWithRole, 0),
		Users:           make(map[string]*notiontypes.UserWithRole, 0),
		Collections:     make(map[string]*notiontypes.CollectionWithRole, 0),
		CollectionViews: make(map[string]*notiontypes.CollectionViewWithRole, 0),
	}
	// TODO: consider merging into first recordmap as a heap optimization.

	for _, rm := range rms {
		for k, v := range rm.Blocks {
			result.Blocks[k] = v
		}
		for k, v := range rm.Space {
			result.Space[k] = v
		}
		for k, v := range rm.Users {
			result.Users[k] = v
		}
		for k, v := range rm.Collections {
			result.Collections[k] = v
		}
		for k, v := range rm.CollectionViews {
			result.CollectionViews[k] = v
		}
	}
	return result, nil
}

func (c *Client) parsePageFromRecordMaps(pageID string, responses []notiontypes.RecordMap) (*Page, error) {
	rm, err := mergeRecordMaps(responses...)
	if err != nil {
		return nil, err
	}
	pageBlock, ok := rm.Blocks[pageID]
	if !ok {
		return nil, fmt.Errorf("notion: missing page id in block list")
	}
	page := &Page{Block: pageBlock.Value}
	blocks := make(map[string]*notiontypes.Block, len(rm.Blocks))
	for k, v := range rm.Blocks {
		blocks[k] = v.Value
	}
	if err := notiontypes.ResolveBlock(page.Block, blocks); err != nil {
		return nil, errors.Wrap(err, "resolveBlock failed")
	}
	return page, nil
}
