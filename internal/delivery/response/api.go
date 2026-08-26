package response

type ApiErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ApiListMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type ApiDataResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type ApiListResponse struct {
	Success bool  `json:"success"`
	Data    []any `json:"data"`
	Meta    ApiListMeta
}

type ApiErrorResponse struct {
	Success bool           `json:"success"`
	Error   ApiErrorDetail `json:"error"`
}

type ApiEmptyResponse struct {
	Success bool `json:"success"`
}
