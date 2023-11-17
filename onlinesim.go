package gonlinesim

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

const (
	defaultHost = "https://onlinesim.io"

	LangFR = "fr"
	LangDE = "de"
	LangRU = "ru"
	LangEN = "en"
	LangZH = "zh"
)

var (
	ErrAccountBlocked    = errors.New("account blocked")
	ErrWrongKey          = errors.New("wrong key")
	ErrNoKey             = errors.New("no key")
	ErrNoService         = errors.New("no service or invalid service name")
	ErrRequestNotFound   = errors.New("request not found")
	ErrAPIAccessDisabled = errors.New("api access disabled")
	ErrAPIAccessIP       = errors.New("api access blocked for this IP")
	ErrLowBalance        = errors.New("low balance")
)

func (c Client) GetFreeList(ctx context.Context,
	opts ...Option) (*GetFreeListResponse, error) {
	response := GetFreeListResponse{}
	req := c.client.R().SetResult(&response).SetContext(ctx)
	for _, opt := range opts {
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
