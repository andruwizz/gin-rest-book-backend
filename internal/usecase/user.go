package usecase

import (
	"time"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Create(req *request.UserCreate) (*entity.User, error)
	Login(req *request.UserLogin) (*entity.UserAuth, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}

func (c *userUsecase) Create(req *request.UserCreate) (*entity.User, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}

	err = c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (c *userUsecase) Login(req *request.UserLogin) (*entity.UserAuth, error) {

	user := new(entity.User)
	err := c.repository.Find(req.Email, user)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, err
	}

	claims := entity.AuthClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "APP_NAME",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1)),
		},
		Name:  user.Name,
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("SIGNATURE_KEY"))
	if err != nil {
		return nil, err
	}

	res := &entity.UserAuth{
		Token: signedToken,
	}

	return res, nil
}
