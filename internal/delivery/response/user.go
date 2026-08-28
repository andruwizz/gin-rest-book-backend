package response

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/gin-gonic/gin"
)

func UserResource(ctx *gin.Context, code int, e *entity.User) {
	res := DataResponse{
		Success: true,
		Data:    e,
	}

	ctx.JSON(code, res)
}

func UserAuthResource(ctx *gin.Context, code int, e *entity.UserAuth) {
	res := DataResponse{
		Success: true,
		Data:    e,
	}

	ctx.JSON(code, res)
}
