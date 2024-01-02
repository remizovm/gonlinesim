package gonlinesim

import (
	"context"
	"fmt"
	"net/http"
)

// GetFreeList depending on the parameters returns information on available
// countries with free numbers, received messages and the list of restricted
// services.
//
// Without parameters: returns a list of available countries with free numbers,
// messages received by the first number from the list and the list of
// restricted services.
//
// With country parameter: returns a list of available free numbers of that
// country, messages received by the first number from the list and the list of
// restricted services.
//
// With country and number parameters: returns a list of SMS messages received
// by this number and the list of restrices services.
// You can send this request to get a number that can be used for SMS receiption
//
//	later on, for example for testing.
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
