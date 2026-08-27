package response

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorDetail struct {
	Code    string `json:"code"`
	Message any    `json:"message"`
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

type ValidationError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

func EmptyResource(ctx *gin.Context, code int) {
	res := EmptyResponse{
		Success: true,
	}

	ctx.JSON(code, res)
}

func ErrorValidation(ctx *gin.Context, code int, errorCode string, err error) {
	res := ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    errorCode,
			Message: ExtractValidationError(err.(validator.ValidationErrors)),
		},
	}

	ctx.JSON(code, res)
}

func ErrorResource(ctx *gin.Context, code int, errorCode string, err error) {
	res := ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    errorCode,
			Message: err.Error(),
		},
	}

	ctx.JSON(code, res)
}

func ExtractValidationError(ve validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}

	for _, f := range ve {
		err := f.Tag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, ValidationError{Field: strings.ToLower(f.Field()), Rule: err})
	}

	return errs
}
