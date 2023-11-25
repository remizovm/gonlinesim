package gonlinesim

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type Client struct {
	client *resty.Client
	apiKey string
}

type timeDecoder struct{}

// Decode implements jsoniter.ValDecoder.
func (timeDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	val := iter.ReadString()
	data, err := time.Parse("2006-01-02 15:04:05", val)
	if err != nil {
		iter.Error = err
		return
	}
	*((*time.Time)(ptr)) = data
}

func NewClient(apiKey string) *Client {
	jsoniter.RegisterTypeDecoder("time.Time", timeDecoder{})
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	result := Client{
		client: resty.New().
			SetJSONMarshaler(json.Marshal).
			SetJSONUnmarshaler(json.Unmarshal).
			OnAfterResponse(checkErrMiddleware).
			SetBaseURL(defaultHost),
		apiKey: apiKey,
	}
	return &result
}

func checkErrMiddleware(c *resty.Client, resp *resty.Response) error {
	response := BaseResp{}
	if err := jsoniter.Unmarshal(resp.Body(), &response); err != nil {
		return err
	}
	if v, ok := response.Response.(string); ok {
		return checkErr(v)
	} else if v, ok := response.Response.(float64); ok {
		return checkErr(fmt.Sprintf("%d", int(v)))
	}
	return nil
}
