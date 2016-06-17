package client

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Client struct
type Client struct {
	c *http.Client
}

// NewClient returns Client
func NewClient() *Client {
	return &Client{c: &http.Client{}}
}

// Hello accesses /hello and return the response body in the string type
func (c *Client) Hello(host, name string) (int, string, error) {
	req, err := http.NewRequest("GET", host+"/hello", nil)
	if err != nil {
		return 0, "", errors.Wrap(err, "failed to create request")
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return resp.StatusCode, "", errors.Wrap(err, "failed to request")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", errors.Wrap(err, "failed to read body")
	}
	return resp.StatusCode, string(body), nil
}

// Goodbye accesses /goodbye and return the response body in the string type
func (c *Client) Goodbye(host, name string) (int, string, error) {
	return 0, "", nil
}
