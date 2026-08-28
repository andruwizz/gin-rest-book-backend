package usecase

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/entity"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
)

type UserUsecase interface {
	Create(req *request.UserCreate) (*entity.User, error)
	Login(req *request.UserLogin) (*entity.UserAuth, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUseUsecase(repository repository.UserRepository) UserUsecase {
	return &userUsecase{repository}
}

func (c *userUsecase) Create(req *request.UserCreate) (*entity.User, error) {
	payload := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := c.repository.Create(payload)
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

	res := &entity.UserAuth{
		User:  *user,
		Token: "secret",
	}

	return res, nil
}
