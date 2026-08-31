package helper

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func EncryptAuthPassword(plain string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(plain), 10)
	if err != nil {
		return "", err
	}

	return string(encrypted), nil
}

func VerifyAuthPassword(plain string, encrypted string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(plain), []byte(encrypted)); err != nil {
		return err
	}

	return nil
}

func EncodeAuthToken(user *entity.User) (*entity.AuthToken, error) {
	duration, _ := strconv.Atoi(os.Getenv("AUTH_TOKEN_DURATION"))
	claims := entity.AuthClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APP_NAME",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Hour)),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("AUTH_SIGNATURE_KEY")))
	if err != nil {
		return nil, err
	}

	resToken := &entity.AuthToken{
		Token: signedToken,
	}

	return resToken, nil
}

func DecodeAuthToken(header string) (*entity.AuthClaim, error) {
	if !strings.Contains(header, "Bearer ") {
		return nil, errors.New("authentication token not found")
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entity.AuthClaim{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("AUTH_SIGNATURE_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*entity.AuthClaim); ok {
		return claims, nil
	}

	return nil, errors.New("failed decoding token")
}

func GetAuthUser(ctx *gin.Context, user *request.UserGet) error {
	userInfo, ok := ctx.Get(os.Getenv("AUTH_CONTEXT_KEY"))
	if !ok {
		return errors.New("Unauthenticated")
	}

	userInfoMap := userInfo.(map[string]string)
	user.Name = userInfoMap["user"]
	user.Email = userInfoMap["email"]

	return nil
}
