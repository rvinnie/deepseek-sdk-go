package deepseek

import (
	"context"
	"encoding/json"
	"github.com/rvinnie/deepseek-sdk-go/deepseek/utils"
	"net/http"
)

const (
	chatEndpointPart        = "chat"
	completionsEndpointPart = "completions"
)

// CreateChatCompletions creates http request for chat completions
func (c *Client) CreateChatCompletions(ctx context.Context, req *ChatCompletionsRequest) (*ChatCompletionsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	reqJson, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		return nil, err
	}

	reqPath := utils.JoinEndpointParts(chatEndpointPart, completionsEndpointPart)
	resp, err := c.makeRequest(ctx, http.MethodPost, reqPath, reqJson)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	parsedResp, err := parseChatCompletionsResponse(resp)
	if err != nil {
		return nil, err
	}

	return parsedResp, nil
}

// CreateChatCompletionsWithStream creates http request for stream chat completions
func (c *Client) CreateChatCompletionsWithStream(ctx context.Context, req *ChatCompletionsRequest) (*http.Response, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	req.Stream = true

	reqJson, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		return nil, err
	}

	reqPath := utils.JoinEndpointParts(chatEndpointPart, completionsEndpointPart)
	resp, err := c.makeRequest(ctx, http.MethodPost, reqPath, reqJson)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseChatCompletionsResponse(resp *http.Response) (*ChatCompletionsResponse, error) {
	responseBytes, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	chatCompletionsResponse := ChatCompletionsResponse{}
	err = json.Unmarshal(responseBytes, &chatCompletionsResponse)
	if err != nil {
		return nil, err
	}

	return &chatCompletionsResponse, nil
}
