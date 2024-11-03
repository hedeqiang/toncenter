package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client defines the interface for making HTTP requests
type Client interface {
	Request(method, endpoint string, body interface{}) ([]byte, error)
}

// HTTPClient implements the Client interface
type HTTPClient struct {
	BaseURL    string
	APIKey     string
	Client     *http.Client
	RetryCount int
	RetryDelay time.Duration
}

// NewClient creates a new HTTP client
func NewClient(baseURL, apiKey string, httpClient *http.Client, retryCount int, retryDelay time.Duration) *HTTPClient {
	return &HTTPClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		Client:     httpClient,
		RetryCount: retryCount,
		RetryDelay: retryDelay,
	}
}

// Request makes an HTTP request
func (c *HTTPClient) Request(method, endpoint string, body interface{}) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Set("X-API-Key", c.APIKey)
	}

	var resp *http.Response
	var respErr error

	for i := 0; i <= c.RetryCount; i++ {
		resp, respErr = c.Client.Do(req)
		if respErr == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if i < c.RetryCount {
			time.Sleep(c.RetryDelay)
		}
	}

	if respErr != nil {
		return nil, fmt.Errorf("do request: %w", respErr)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	return respBody, nil
}
