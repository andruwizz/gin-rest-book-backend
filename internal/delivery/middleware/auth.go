package middleware

import (
	"net/http"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func Authenticated(authHelper *helper.AuthHelper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		claims, err := authHelper.DecodeAuthToken(authHeader)
		if err != nil {
			response.ErrorResource(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		userInfo := map[string]string{"name": claims.Name, "email": claims.Email}
		authHelper.SetAuthContext(ctx, userInfo)
		ctx.Next()
	}
}
