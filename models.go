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

type Link struct {
	URL    string `json:"url"`
	Label  string `json:"label"`
	Active bool   `json:"active"`
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
		CurrentPage  int        `json:"current_page"`
		Data         []*Message `json:"data"`
		FirstPageURL string     `json:"first_page_url"`
		From         int        `json:"from"`
		LastPage     int        `json:"last_page"`
		LastPageURL  string     `json:"last_page_url"`
		Links        []*Link    `json:"links"`
		NextPageURL  string     `json:"next_page_url"`
		Path         string     `json:"path"`
		PerPage      int        `json:"per_page"`
		PrevPageURL  string     `json:"prev_page_url"`
		To           int        `json:"to"`
		Total        int        `json:"total"`
		Number       string     `json:"number"`
		Country      int        `json:"country"`
	} `json:"messages"`
	Ignore string `json:"ignore"`
}
