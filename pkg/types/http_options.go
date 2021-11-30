package types

import (
	"net/http"
	neturl "net/url"
)

type TinderboxHTTPOptions struct {
	retries      int         // Max retries for an HTTP request
	backoffDelay int         // Delay (in seconds) to backoff on each retry
	target       string      // Target for an HTTP request
	headers      http.Header // Headers for an HTTP request
}

// NewTinderboxHTTPOptions returns an empty instance of TinderboxHTTPOptions
func NewTinderboxHTTPOptions() *TinderboxHTTPOptions {
	return &TinderboxHTTPOptions{
		retries:      3,
		backoffDelay: 1,
		target:       "",
		headers:      http.Header{},
	}
}

// WithMaxRetries sets the maximum number of HTTP retries for a Tinderbox-based HTTP request
func (h *TinderboxHTTPOptions) WithMaxRetries(i int) *TinderboxHTTPOptions {
	h.retries = i

	return h
}

// WithHeader sets exactly one HTTP header.
// Invoke for each header desired to be set.
func (h *TinderboxHTTPOptions) WithHeader(headerName, headerVal string) *TinderboxHTTPOptions {
	h.headers.Set(headerName, headerVal)

	return h
}

// WithDelay sets the desired delay factor f (f * 1 second)
func (h *TinderboxHTTPOptions) WithDelay(f int) *TinderboxHTTPOptions {
	h.backoffDelay = f

	return h
}

// ToRequest returns an HTTP request object along with any errors encountered during request generation.
func (h *TinderboxHTTPOptions) ToRequest(url string) (*http.Request, error) {
	h.target = url
	pURL, err := neturl.Parse(url)
	if err != nil {
		return &http.Request{}, err
	}

	return &http.Request{
		Method: http.MethodGet,
		URL:    pURL,
		Header: h.headers,
	}, nil
}

func (h *TinderboxHTTPOptions) Retries() int {
	return h.retries
}

func (h *TinderboxHTTPOptions) BackoffDelay() int {
	return h.backoffDelay
}
