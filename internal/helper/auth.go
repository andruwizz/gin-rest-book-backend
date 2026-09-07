package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper struct {
	appName       string
	contextKey    string
	tokenDuration int // in hours
	signatureKey  string
}

func NewAuthHelper(appName string, contextKey string, tokenDuration int, signatureKey string) *AuthHelper {
	return &AuthHelper{appName, contextKey, tokenDuration, signatureKey}
}

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

func (h *AuthHelper) EncodeAuthToken(user *entity.User) (*entity.AuthToken, error) {
	duration := h.tokenDuration
	claims := entity.AuthClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    h.appName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Hour)),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(h.signatureKey))
	if err != nil {
		return nil, err
	}

	resToken := &entity.AuthToken{
		Token: signedToken,
	}

	return resToken, nil
}

func (h *AuthHelper) DecodeAuthToken(header string) (*entity.AuthClaim, error) {
	if !strings.Contains(header, "Bearer ") {
		return nil, errors.New("authentication token not found")
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, &entity.AuthClaim{}, func(t *jwt.Token) (any, error) {
		return []byte(h.signatureKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*entity.AuthClaim); ok {
		return claims, nil
	}

	return nil, errors.New("failed decoding token")
}

func (h *AuthHelper) SetAuthContext(ctx *gin.Context, userInfo map[string]string) {
	ctx.Set(h.contextKey, userInfo)
}

func (h *AuthHelper) GetAuthUser(ctx *gin.Context, user *entity.User) error {
	userInfo, ok := ctx.Get(h.contextKey)
	if !ok {
		return errors.New("Unauthenticated")
	}

	userInfoMap := userInfo.(map[string]string)
	user.Name = userInfoMap["name"]
	user.Email = userInfoMap["email"]

	return nil
}
