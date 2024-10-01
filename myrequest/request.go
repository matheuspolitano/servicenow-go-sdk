package myrequest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// HTTPClient is a struct that wrap an http.Client with additional configuration
type HTTPClient struct {
	client *http.Client
	baseURL string
	headers map[string]string
	username string
	password string
}


// NewHTTPClient creates and return a new HTTPClient with custom settings
func NewHTTPClient(baseURL string, timeout time.Duration, headers map[string]string, username string, password string) *HTTPClient{
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
		headers: headers,
		username:  username,
		password: password,
	}
}

// Get performs a GET request to the specified endpoint with context for timeout control
func (h *HTTPClient) Get(ctx context.Context, endpoint string) (map[string]any, error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.baseURL + endpoint, nil)
	if err != nil{
		return nil, err
	}
	return h.makeRequest(req)
}

func (h *HTTPClient) makeRequest(req *http.Request)(map[string]any, error){
	h.setHeaders(req)

	if h.username != "" && h.password != ""{
		req.SetBasicAuth(h.username, h.password)
	}

	resp, err := h.client.Do(req)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
		log.Info().Str("Method", req.Method).Str("path", req.URL.Path).Str("query", req.URL.RawQuery).Int("Status", resp.StatusCode).Msg("Request")
	if resp.StatusCode >= http.StatusBadRequest{
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	
	const maxBodySize = 10 * 1024 * 1024
	bbody, err := io.ReadAll(io.LimitReader(resp.Body, maxBodySize))
	if err != nil{
		return nil, err
	}
	var body map[string]any
	if err = json.Unmarshal(bbody, &body); err != nil{
		return nil, err
	}
	
	return body, nil
}

// Post performs a POST request to the specified endpoint and body with context for timeout control
func (h *HTTPClient) Post(ctx context.Context, endpoint string, body any) (map[string]any, error){
	bodyRequest, err := json.Marshal(body)
	if err != nil{
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, h.baseURL + endpoint, bytes.NewReader(bodyRequest))
	if err != nil{
		return nil, err
	}
	return h.makeRequest(req)
}


// setHeaders set the HTTPClient headers into the http.Request
func (h *HTTPClient)setHeaders(req *http.Request){
	for key, value := range h.headers{
		req.Header.Set(key, value)
	}

}