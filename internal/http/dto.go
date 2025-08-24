package http

// ListQuery represents the query parameters for listing resources
type ListQuery struct {
	Query    string `form:"query"`
	SortBy   string `form:"sort"`
	Order    string `form:"order"`
	Page     int    `form:"page"`
	PageSize int    `form:"page_size"`
}

// ListResponse represents a paginated response for listing resources
type ListResponse[T any] struct {
	Data []T `json:"data"`
	Meta struct {
		Page     int   `json:"page"`
		PageSize int   `json:"page_size"`
		Total    int64 `json:"total"`
	} `json:"meta"`
}
