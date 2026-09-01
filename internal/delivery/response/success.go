package response

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/gin-gonic/gin"
)

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

type EmptyResponse struct {
	Success bool `json:"success"`
}

type Data interface {
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

func Resource[T Data](ctx *gin.Context, code int, e *T) {
	res := DataResponse{
		Success: true,
		Data:    e,
	}

	ctx.JSON(code, res)
}

func Collection[T Data](ctx *gin.Context, code int, e []T, p *entity.Pagination) {
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
