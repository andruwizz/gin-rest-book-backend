package response

import (
	"github.com/gin-gonic/gin"
)

// Standard Response Wrapper
type Response struct {
	Success bool         `json:"success"`
	Data    any          `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
	Meta    *Meta        `json:"meta,omitempty"`
}

type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// OK Success response
func Success(ctx *gin.Context, status int, data any, meta any) {
	var response = Response{
		Success: true,
	}

	if data != nil {
		response.Data = data
	}

	if meta != nil {
		response.Meta = &Meta{
			Page:       meta.(Meta).Page,
			PerPage:    meta.(Meta).PerPage,
			Total:      meta.(Meta).Total,
			TotalPages: meta.(Meta).TotalPages,
		}
	}

	ctx.JSON(status, response)
}

// Fail Error response
func Fail(ctx *gin.Context, status int, code, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error:   &ErrorDetail{Code: code, Message: message},
	})
}
