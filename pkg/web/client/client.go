package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/loikx/lazylab/pkg/web/endpoint"
	"github.com/loikx/lazylab/pkg/web/headers"
	"github.com/loikx/lazylab/pkg/web/query"
	"github.com/loikx/lazylab/pkg/web/request"
)

var ErrNoHttpClient = errors.New("http client is nil")

type (
	HeaderKey   string
	HeaderValue string
)

type Client struct {
	retries int
	client  *http.Client
	request *request.Request
}

func New(retries int, client *http.Client) *Client {
	return &Client{
		retries: retries,
		client:  client,
		request: nil,
	}
}

func (c *Client) Do() (*http.Response, error) {
	if c.client == nil {
		return nil, fmt.Errorf("can not perform request: %w", ErrNoHttpClient)
	}

	response, err := c.client.Do(c.request.ToHttpRequest())
	if err != nil {
		return response, fmt.Errorf("error when doing client request: %w", err)
	}

	return response, nil
}

func (c *Client) WithEndpoint(endpointURL string) error {
	_, err := url.ParseRequestURI(endpointURL)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}

	c.request.WithEndpoint(endpoint.FromString(endpointURL))
	return nil
}

func (c *Client) WithHeaders(rawHeaders map[string]string) {
	headers := headers.New(rawHeaders)
	c.request.WithHeaders(headers)
}

func (c *Client) WithQuery(rawQuery map[string]string) {
	queries := query.New(rawQuery)
	c.request.WithQuery(queries)
}

func (c *Client) Request() string {
	return c.request.ToURL()
}
