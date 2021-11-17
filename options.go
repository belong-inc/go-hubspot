package hubspot

import (
	"net/http"
	"net/url"
)

type Option func(c *Client)

// TODO: Add WithRetryConfig

func WithAPIVersion(version string) Option {
	return func(c *Client) {
		c.apiVersion = version
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.HTTPClient = client
	}
}

func WithBaseURL(url *url.URL) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}
