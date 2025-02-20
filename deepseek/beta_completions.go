package deepseek

import (
	"context"
	"encoding/json"
	"github.com/rvinnie/deepseek-sdk-go/deepseek/utils"
	"net/http"
)

const (
	betaChatEndpointPart        = "beta"
	betaCompletionsEndpointPart = "completions"
)

// CreateBetaCompletions creates http request for FIM completions
func (c *Client) CreateBetaCompletions(ctx context.Context, req *BetaCompletionsRequest) (*BetaCompletionsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	reqJson, err := json.MarshalIndent(req, "", "\t")
	if err != nil {
		return nil, err
	}

	reqPath := utils.JoinEndpointParts(betaChatEndpointPart, betaCompletionsEndpointPart)
	resp, err := c.makeRequest(ctx, http.MethodPost, reqPath, reqJson)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	parsedResp, err := parseBetaCompletionsResponse(resp)
	if err != nil {
		return nil, err
	}

	return parsedResp, nil
}

func parseBetaCompletionsResponse(resp *http.Response) (*BetaCompletionsResponse, error) {
	responseBytes, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	chatCompletionsResponse := BetaCompletionsResponse{}
	err = json.Unmarshal(responseBytes, &chatCompletionsResponse)
	if err != nil {
		return nil, err
	}

	return &chatCompletionsResponse, nil
}
