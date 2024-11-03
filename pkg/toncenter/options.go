package toncenter

import "time"

// Option is a function that configures a Client
type Option func(*Client)

// WithBaseURL sets the base URL for the client
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithTimeout sets the timeout for HTTP requests
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithRetryCount sets the number of retries for failed requests
func WithRetryCount(count int) Option {
	return func(c *Client) {
		c.retryCount = count
	}
}
