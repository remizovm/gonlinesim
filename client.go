package gonlinesim

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	result := Client{
		client: resty.New(),
		apiKey: apiKey,
	}
	result.client.SetBaseURL(defaultHost)
	result.client = result.client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		response := BaseResp{}
		if err := json.Unmarshal(resp.Body(), &response); err != nil {
			return err
		}
		return checkErr(response.Response)
	})
	return &result
}
