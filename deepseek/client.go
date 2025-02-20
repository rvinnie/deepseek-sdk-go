package deepseek

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL               = "https://api.deepseek.com/"
	DefaultTimeoutSeconds = 120
)

type Client struct {
	client *http.Client
	config Config
}

func NewClient(apiKey string) (*Client, error) {
	config := Config{
		apiKey:         apiKey,
		timeoutSeconds: DefaultTimeoutSeconds,
	}

	return NewClientWithConfig(config)
}

func NewClientWithConfig(config Config) (*Client, error) {
	if config.apiKey == "" {
		return nil, errors.New("config error: api key is empty")
	}
	if config.timeoutSeconds <= 0 {
		return nil, errors.New("config error: timeout seconds is invalid")
	}

	return &Client{
		config: config,
		client: &http.Client{
			Timeout: time.Duration(config.timeoutSeconds) * time.Second,
		},
	}, nil
}

func (c *Client) makeRequest(
	ctx context.Context,
	method string,
	endpoint string,
	body []byte,
) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", BaseURL, endpoint)

	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.config.apiKey))

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseResponse(resp *http.Response) ([]byte, error) {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, getError(resp.Body)
	}

	var responseBytes []byte
	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}
