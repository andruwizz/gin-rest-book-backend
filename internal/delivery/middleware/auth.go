package middleware

import (
	"net/http"
	"os"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		claims, err := helper.DecodeAuthToken(authHeader)
		if err != nil {
			response.ErrorResource(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		userInfo := map[string]string{"name": claims.Name, "email": claims.Email}
		ctx.Set(os.Getenv("AUTH_CONTEXT_KEY"), userInfo)
		ctx.Next()
	}
}
