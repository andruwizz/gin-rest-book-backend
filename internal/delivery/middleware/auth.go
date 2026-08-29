package middleware

import (
	"net/http"
	"strings"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type unauthorized struct {
	message string
}

func (u *unauthorized) Error() string {
	return u.message
}

func Authenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer ") {
			err := &unauthorized{message: "Unauthorized"}
			response.ErrorResource(ctx, http.StatusUnauthorized, err)
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &entity.AuthClaim{}, func(t *jwt.Token) (any, error) {
			return []byte("SIGNATURE_KEY"), nil
		})

		if err != nil || !token.Valid {
			err = &unauthorized{message: "Unauthorized"}
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
