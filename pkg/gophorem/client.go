package gophorem

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client makes all the API calls to dev.to.
type Client struct {
	apiKey  string
	baseURL string
	http    *http.Client
}

// Arguments are used for passing query parameters to the dev.to api.
type Arguments map[string]string

// Defaults returns an empty map of arguments.
func Defaults() Arguments {
	return make(map[string]string)
}

func (a Arguments) toQueryParams() url.Values {
	res := make(url.Values)
	for k, v := range a {
		res.Add(k, v)
	}
	return res
}

// Option allows the client to be configured with different options.
type Option func(*Client)

func withBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithAPIKey sets the dev.to api key to use for this client.
// see https://docs.dev.to/api/#section/Authentication for how to set one up.
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// NewClient creates a dev.to client with the provided options.
func NewClient(foremURL string, opts ...Option) *Client {
	res := &Client{
		baseURL: foremURL,
		http:    &http.Client{},
	}
	for _, o := range opts {
		o(res)
	}

	return res
}

// NewDevtoClient creates a dev.to client with the provided options.
func NewDevtoClient(opts ...Option) *Client {
	res := &Client{
		baseURL: "https://dev.to/api",
		http:    &http.Client{},
	}
	for _, o := range opts {
		o(res)
	}

	return res
}

func (c *Client) getRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {

	b := bytes.NewBuffer(nil)
	if method == http.MethodPost || method == http.MethodPut {
		j, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		b = bytes.NewBuffer(j)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, b)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Add("api-key", c.apiKey)
	}

	return req, err
}

// get returns an error if the http client cannot perform a HTTP GET for the provided URL.
func (c *Client) get(ctx context.Context, url string, target interface{}) error {
	req, err := c.getRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error from forem api")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &target)
}

// save returns an error if the http client cannot save the request to dev.to..
func (c *Client) save(ctx context.Context, httpMethod string, url string, payload interface{}, target interface{}) error {
	req, err := c.getRequest(ctx, httpMethod, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("error from forem api. httpCode: %d, response: %s", resp.StatusCode, b)
	}

	return json.Unmarshal(b, &target)
}

// put returns an error if the http client cannot perform a HTTP PUT for the provided URL.
func (c *Client) put(ctx context.Context, url string, payload interface{}, target interface{}) error {
	return c.save(ctx, http.MethodPut, url, payload, target)
}

// post returns an error if the http client cannot perform a HTTP POST for the provided URL.
func (c *Client) post(ctx context.Context, url string, payload interface{}, target interface{}) error {
	return c.save(ctx, http.MethodPost, url, payload, target)
}

// delete returns an error if the http client cannot perform a HTTP DELETE for the provided URL.
func (c *Client) delete(ctx context.Context, url string, payload interface{}) error {
	req, err := c.getRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("error from forem api")
	}

	return nil
}
