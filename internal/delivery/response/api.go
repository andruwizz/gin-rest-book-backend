package response

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ListMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type DataResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

type ListResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
	Meta    ListMeta
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   ErrorDetail `json:"error"`
}

type EmptyResponse struct {
	Success bool `json:"success"`
}
