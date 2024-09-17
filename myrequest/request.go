package myrequest

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

// HTTPClient is a struct that wrap an http.Client with additional configuration
type HTTPClient struct {
	client *http.Client
	baseURL string
	headers map[string]string
}


// NewHTTPClient creates and return a new HTTPClient with custom settings
func NewHTTPClient(baseURL string, timeout time.Duration, headers map[string]string) *HTTPClient{
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		headers: headers,
	}
}

// Get performs a GET request to the specified endpoint with context for timeout control
func (h *HTTPClient) Get(ctx context.Context, endpoint string) (*http.Response, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.baseURL + endpoint, nil)
	if err != nil{
		return nil, err
	}

	h.setHeaders(req)

	resp, err := h.client.Do(req)
	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusOK{
		return nil, errors.New("unexpected status code:" + resp.Status)
	}
	return resp, nil
}

// Post performs a POST request to the specified endpoint and body with context for timeout control
func (h *HTTPClient) Post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.baseURL + endpoint, body)
	if err != nil{
		return nil, err
	}

	h.setHeaders(req)

	resp, err := h.client.Do(req)
	if err != nil{
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated{
		return nil, errors.New("unexpected status code: " + resp.Status) 
	}

	return resp, nil
}


// setHeaders set the HTTPClient headers into the http.Request
func (h *HTTPClient)setHeaders(req *http.Request){
	for key, value := range h.headers{
		req.Header.Set(key, value)
	}

}