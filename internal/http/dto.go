package http

type ListQuery struct {
	Query    string `form:"query"`
	SortBy   string `form:"sort"`
	Order    string `form:"order"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

type ListResponse[T any] struct {
	Data []T `json:"data"`
	Meta struct {
		Page     int   `json:"page"`
		PageSize int   `json:"page_size"`
		Total    int64 `json:"total"`
	} `json:"meta"`
}
