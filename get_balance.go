package gonlinesim

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetBalance(ctx context.Context,
	opts ...Option) (*GetBalanceResp, error) {
	response := GetBalanceResp{}
	req := c.client.R().SetResult(&response).
		SetContext(ctx).SetQueryParam("apikey", c.apiKey)
	for _, opt := range opts {
		opt(req)
	}
	resp, err := req.Get("/api/getBalance.php")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", resp.StatusCode())
	}
	return &response, nil
}
