package helper

import (
	"time"

	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
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

func CreateAuthToken(user *entity.User) (*entity.AuthToken, error) {
	claims := entity.AuthClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APP_NAME",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("SIGNATURE_KEY"))
	if err != nil {
		return nil, err
	}

	resToken := &entity.AuthToken{
		Token: signedToken,
	}

	return resToken, nil
}
