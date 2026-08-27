package response

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/gin-gonic/gin"
)

func BookResource(ctx *gin.Context, code int, e *entity.Book) {
	res := DataResponse{
		Success: true,
		Data:    e,
	}

	ctx.JSON(code, res)
}

func BookCollection(ctx *gin.Context, code int, e []entity.Book, p *entity.Pagination) {
	col := ListResponse{
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

func EmptyResource(ctx *gin.Context, code int) {
	res := EmptyResponse{
		Success: true,
	}

	ctx.JSON(code, res)
}

func ErrorResource(ctx *gin.Context, code int, errorCode, message string) {
	res := ErrorResponse{
		Success: false,
		Error: ErrorDetail{
			Code:    errorCode,
			Message: message,
		},
	}

	ctx.JSON(code, res)
}
