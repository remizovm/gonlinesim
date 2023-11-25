package gonlinesim

import (
	"context"
	"fmt"
	"net/http"
)

func (c *Client) GetTariffs(ctx context.Context,
	opts ...Option) (*GetTariffsResp, error) {
	response := GetTariffsResp{}
	req := c.client.R().SetResult(&response).
		SetContext(ctx).SetQueryParam("apikey", c.apiKey)
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(req)
	}
	resp, err := req.Get("/api/getTariffs.php")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", resp.StatusCode())
	}
	return &response, nil
}
