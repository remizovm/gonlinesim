package gonlinesim

import "time"

type BaseResp struct {
	Response interface{} `json:"response"`
}

type GetTariffsResp struct {
	BaseResp
	Country int    `json:"country"`
	Page    int    `json:"page"`
	End     bool   `json:"end"`
	Filter  string `json:"filter"`
}

type GetBalanceResp struct {
	BaseResp
	Balance  string `json:"balance"`
	ZBalance string `json:"zbalance"`
}

type Message struct {
	Text              string    `json:"text"`
	InNumber          string    `json:"in_number"`
	MyNumber          float64   `json:"my_number"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedAtReadable string    `json:"data_humans"`
}

type GetFreeListResponse struct {
	BaseResp
	Countries []struct {
		PhoneCode    int    `json:"country"`
		Name         string `json:"country_text"`
		OriginalName string `json:"country_original"`
	} `json:"countries"`
	Numbers map[string]struct {
		CountryCode int    `json:"country"`
		CountryName string `json:"country_original"`
		DateHumans  string `json:"data_humans"`
		FullNumber  string `json:"full_number"`
		Archived    bool   `json:"is_archive"`
	} `json:"numbers"`
	Messages struct {
		CurrentPage int        `json:"current_page"`
		Data        []*Message `json:"data"`
	} `json:"messages"`
}
