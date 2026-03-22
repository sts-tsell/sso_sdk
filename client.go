// Package ssosdk provides a Go client for the SSO API.
package ssosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the SSO API client.
type Client struct {
	baseURL     string
	bearerToken string
	httpClient  *http.Client
}

// Option is a functional option for Client.
type Option func(*Client)

// WithBearerToken sets the Authorization Bearer token on the client.
func WithBearerToken(token string) Option {
	return func(c *Client) {
		c.bearerToken = token
	}
}

// WithHTTPClient replaces the default HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// New creates a new Client with the given base URL and optional options.
//
// Example:
//
//	client := ssosdk.New("https://api.example.com", ssosdk.WithBearerToken("my-token"))
func New(baseURL string, opts ...Option) *Client {
	c := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// SetBearerToken updates the bearer token used for authenticated requests.
func (c *Client) SetBearerToken(token string) {
	c.bearerToken = token
}

// do executes an HTTP request and decodes the JSON response into v.
// If v is nil the response body is discarded.
func (c *Client) do(method, path string, body, v interface{}, authenticated bool) error {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("ssosdk: marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("ssosdk: create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	if authenticated {
		if c.bearerToken == "" {
			return ErrNoToken
		}
		req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("ssosdk: execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return parseErrorResponse(resp)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("ssosdk: decode response: %w", err)
		}
	}

	return nil
}
