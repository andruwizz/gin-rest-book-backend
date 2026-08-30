package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/helper"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type UserUsecase interface {
	Create(req *request.UserCreate) (*entity.User, error)
	Login(req *request.UserLogin) (*entity.AuthToken, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}

func (c *userUsecase) Create(req *request.UserCreate) (*entity.User, error) {
	password, err := helper.EncryptAuthPassword(req.Password)
	if err != nil {
		return nil, err
	}

	payload := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: password,
	}

	err = c.repository.Create(payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (c *userUsecase) Login(req *request.UserLogin) (*entity.AuthToken, error) {

	user := new(entity.User)
	err := c.repository.Find(req.Email, user)
	if err != nil {
		return nil, err
	}

	if err := helper.VerifyAuthPassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	res, err := helper.EncodeAuthToken(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}
