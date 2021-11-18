package dexcomClient

import "fmt"

type Client struct {
	token   string
	sandbox bool
}

func (c *Client) getURL(path string) string {
	base := baseUrl
	if c.sandbox {
		base = sandboxUrl
	}
	return fmt.Sprintf("%s%s", base, path)
}

func NewClient(token string) *Client {
	dc := &Client{
		token: token,
	}
	return dc
}
