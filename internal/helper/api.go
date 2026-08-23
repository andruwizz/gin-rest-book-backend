package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Standard Response Wrapper
type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
	Meta    *Meta      `json:"meta,omitempty"`
}

type ErrorInfo struct {
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
func OK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func OkWithMeta(ctx *gin.Context, data any, meta Meta) {
	ctx.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
		Meta: &Meta{
			Page:       meta.Page,
			PerPage:    meta.PerPage,
			Total:      meta.Total,
			TotalPages: meta.TotalPages,
		},
	})
}

// Fail Error response
func Fail(ctx *gin.Context, status int, code, message string) {
	ctx.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}
