package gonlinesim

import (
	"context"
	"fmt"
	"net/http"
)

func (c Client) GetFreeList(ctx context.Context,
	opts ...Option) (*GetFreeListResponse, error) {
	response := GetFreeListResponse{}
	req := c.client.R().SetResult(&response).SetContext(ctx)
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(req)
	}
	resp, err := req.Get("/api/getFreeList")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("bad status code %d", resp.StatusCode())
	}
	return &response, nil
}
