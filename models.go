package gonlinesim

type BaseResp struct {
	Result string `json:"result"`
}

type GetBalanceResp struct {
	BaseResp
	Response string `json:"response"`
	Balance  string `json:"balance"`
	ZBalance string `json:"zbalance"`
}

type GetFreeListResponse struct {
	Countries []struct {
		PhoneCode int    `json:"country"`
		Name      string `json:"country_text"`
	} `json:"countries"`
	Numbers  interface{} `json:"numbers"`
	Messages struct {
		CurrentPage int `json:"current_page"`
		Data        []struct {
			Text              string      `json:"text"`
			InNumber          string      `json:"in_number"`
			MyNumber          interface{} `json:"my_number"`
			CreatedAt         string/*time.Time*/ `json:"created_at"`
			CreatedAtReadable string `json:"data_humans"`
		} `json:"data"`
	} `json:"messages"`
}
