package middleware

import (
	"net/http"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		token, err := helper.DecodeAuthToken(authHeader)
		if err != nil {
			response.ErrorResource(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(entity.AuthClaim); ok {
			ctx.Set("userInfo", map[string]string{"name": claims.Name, "email": claims.Email})
			ctx.Next()
		}
	}
}
