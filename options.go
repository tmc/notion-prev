package notion

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// ClientOption allows customization of Clients.
type ClientOption func(*Client)

// WithBaseURL allows configuration on of a custom base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithToken allows configuration on of an authentication token.
func WithToken(token string) ClientOption {
	return func(c *Client) {
		c.token = token
	}
}

// WithHTTPClient allows customization of the http.Client that is used for API communication.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithLogger allows configuration of the Logger.
//
// See the WrapLogrus utility type to supply a logrus Logger.
func WithLogger(logger Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

// WithDebugLogging attaches a debug-level logger to the client.
var WithDebugLogging ClientOption = func(c *Client) {
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	c.logger = &WrapLogrus{logger}
}
