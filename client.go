package creem

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	key  string
	http *http.Client
	test bool
}

func New(key string, opts ...Option) *Client {
	ret := &Client{
		key: key,
	}
	for _, opt := range opts {
		opt(ret)
	}
	if ret.http == nil {
		ret.http = http.DefaultClient
	}
	return ret
}

func (c *Client) SetHTTPClient(clt *http.Client) *Client {
	c.http = clt
	return c
}

func (c *Client) SetTestMode(v bool) *Client {
	c.test = v
	return c
}

func (c *Client) BaseURL() string {
	if c.test {
		return TestBaseURL
	}
	return BaseURL
}

func (c *Client) Do(ctx context.Context, req Request, resp any) error {
	var buf io.ReadWriter
	if req.Method() == http.MethodPost {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(req); err != nil {
			return err
		}
	}
	gw := fmt.Sprintf("%s/%s", c.BaseURL(), req.Gateway())
	httpReq, err := http.NewRequestWithContext(ctx, req.Method(), gw, buf)
	if err != nil {
		return err
	}
	return c.fetch(httpReq, resp)
}

func (c *Client) fetch(req *http.Request, resp any) error {
	req.Header.Set("x-api-key", c.key)
	req.Header.Set("Content-type", "application/json")
	httpResp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		switch httpResp.StatusCode {
		case 400:
			return ErrInvalidParameters
		case 401:
			return ErrMissingAPIKey
		case 403:
			return ErrInvalidAPIKey
		case 404:
			return ErrNoResourceFound
		case 429:
			return ErrExceededRateLimit
		case 500:
			return ErrServerError
		default:
			return ErrUnknown
		}
	}
	if resp != nil {
		return json.NewDecoder(httpResp.Body).Decode(resp)
	}
	return nil
}
