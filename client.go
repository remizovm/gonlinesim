package gonlinesim

import "github.com/go-resty/resty/v2"

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
	return &result
}
