package creem

import "net/http"

type Option func(*Client)

func WithHTTPClient(clt *http.Client) Option {
	return func(c *Client) {
		c.http = clt
	}
}

func WithTestMode(v bool) Option {
	return func(c *Client) {
		c.test = v
	}
}
