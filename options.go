package notion

import "net/http"

// ClientOption allows customization of Clients.
type ClientOption func(*Client)

// WithBaseURL allows configuration on of a custom base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient allows customization of the http.Client that is used for API communication.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}
