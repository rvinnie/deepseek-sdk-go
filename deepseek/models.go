package deepseek

import (
	"context"
	"encoding/json"
	"github.com/rvinnie/deepseek-sdk-go/deepseek/utils"
	"net/http"
)

const modelsEndpointPart = "models"

// GetModels creates http request for getting models
func (c *Client) GetModels(ctx context.Context) (*ModelsResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	reqPath := utils.JoinEndpointParts(modelsEndpointPart)
	resp, err := c.makeRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	parsedResp, err := parseModelsResponse(resp)
	if err != nil {
		return nil, err
	}

	return parsedResp, nil
}

func parseModelsResponse(resp *http.Response) (*ModelsResponse, error) {
	responseBytes, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	balanceResponse := ModelsResponse{}
	err = json.Unmarshal(responseBytes, &balanceResponse)
	if err != nil {
		return nil, err
	}

	return &balanceResponse, nil
}
