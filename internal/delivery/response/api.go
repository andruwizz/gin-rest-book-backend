package response

import (
	"fmt"
	"strings"

	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type errorDetail struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type listMeta struct {
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
	Success bool     `json:"success"`
	Data    any      `json:"data"`
	Meta    listMeta `json:"meta"`
}

type ErrorResponse struct {
	Success bool        `json:"success"`
	Error   errorDetail `json:"error"`
}

type EmptyResponse struct {
	Success bool `json:"success"`
}

func EmptyResource(ctx *gin.Context, code int) {
	res := EmptyResponse{
		Success: true,
	}

	ctx.JSON(code, res)
}

func ErrorValidation(ctx *gin.Context, code int, err error) {
	msg, errs := extractValidationError(err.(validator.ValidationErrors))
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: msg,
			Errors:  errs,
		},
	}

	ctx.JSON(code, res)
}

func ErrorResource(ctx *gin.Context, code int, err error) {
	res := ErrorResponse{
		Success: false,
		Error: errorDetail{
			Message: err.Error(),
		},
	}

	ctx.JSON(code, res)
}

func extractValidationError(ve validator.ValidationErrors) (string, map[string]string) {
	errs := map[string]string{}
	for _, f := range ve {
		err := strings.Split(f.Error(), "Error:")
		errs[strings.ToLower(f.Field())] = fmt.Sprintf("%s.", err[1])
	}

	message := fmt.Sprintf("%s.", strings.Split(ve[0].Error(), "Error:")[1])
	if errCount := len(ve); errCount > 1 {
		message = fmt.Sprintf("%s (and %d more)", message, errCount-1)
	}

	return message, errs
}

func Resource[T Entity](ctx *gin.Context, code int, e *T) {
	res := DataResponse{
		Success: true,
		Data:    e,
	}

	ctx.JSON(code, res)
}

func Collection[T Entity](ctx *gin.Context, code int, e []T, p *entity.Pagination) {
	col := ListResponse{
		Success: true,
		Data:    e,
		Meta: listMeta{
			Page:       p.Page,
			PerPage:    p.Limit,
			Total:      int(p.TotalRecords),
			TotalPages: p.TotalPage,
		},
	}

	ctx.JSON(code, col)
}

type Entity interface {
	entity.Book |
		entity.User |
		entity.AuthToken
}
