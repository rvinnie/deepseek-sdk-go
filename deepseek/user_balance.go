package deepseek

import (
	"context"
	"encoding/json"
	"github.com/rvinnie/deepseek-sdk-go/deepseek/utils"
	"net/http"
)

const (
	userEndpointPart    = "user"
	balanceEndpointPart = "balance"
)

// GetUserBalance creates http request for user balance
func (c *Client) GetUserBalance(ctx context.Context) (*BalanceResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	reqPath := utils.JoinEndpointParts(userEndpointPart, balanceEndpointPart)
	resp, err := c.makeRequest(ctx, http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	parsedResp, err := parseUserBalanceResponse(resp)
	if err != nil {
		return nil, err
	}

	return parsedResp, nil
}

func parseUserBalanceResponse(resp *http.Response) (*BalanceResponse, error) {
	responseBytes, err := parseResponse(resp)
	if err != nil {
		return nil, err
	}

	balanceResponse := BalanceResponse{}
	err = json.Unmarshal(responseBytes, &balanceResponse)
	if err != nil {
		return nil, err
	}

	return &balanceResponse, nil
}
