package response

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/gin-gonic/gin"
)

type ListMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type DataResponse[T entityType] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type ListResponse[T entityType] struct {
	Success bool     `json:"success"`
	Data    []T      `json:"data"`
	Meta    ListMeta `json:"meta"`
}

type EmptyResponse struct {
	Success bool `json:"success"`
}

type entityType interface {
	entity.Book |
		entity.User |
		entity.AuthToken
}

func Empty(ctx *gin.Context, code int) {
	res := EmptyResponse{
		Success: true,
	}

	ctx.JSON(code, res)
}

func Resource[T entityType](ctx *gin.Context, code int, e *T) {
	res := DataResponse[T]{
		Success: true,
		Data:    *e,
	}

	ctx.JSON(code, res)
}

func Collection[T entityType](ctx *gin.Context, code int, e []T, p *entity.Pagination) {
	col := ListResponse[T]{
		Success: true,
		Data:    e,
		Meta: ListMeta{
			Page:       p.Page,
			PerPage:    p.Limit,
			Total:      int(p.TotalRecords),
			TotalPages: p.TotalPage,
		},
	}

	ctx.JSON(code, col)
}
