package user

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserUsecase interface {
	Create(req *UserCreateParam) (*entity.User, error)
	Login(req *UserLoginParam) (*entity.AuthToken, error)
	Get(req *UserGetParam) (*entity.User, error)
	GetCurrentUser(ctx *gin.Context, req *UserGetParam) error
}

type userUsecase struct {
	repository repository.UserRepository
	authHelper *helper.AuthHelper
}

func NewUserUsecase(repository repository.UserRepository, authHelper *helper.AuthHelper) UserUsecase {
	return &userUsecase{repository, authHelper}
}

func (c *userUsecase) Create(req *UserCreateParam) (*entity.User, error) {
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

func (c *userUsecase) Login(req *UserLoginParam) (*entity.AuthToken, error) {
	user := new(entity.User)
	err := c.repository.Find(req.Email, user)
	if err != nil {
		return nil, err
	}

	if err := helper.VerifyAuthPassword(user.Password, req.Password); err != nil {
		return nil, err
	}

	res, err := c.authHelper.EncodeAuthToken(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *userUsecase) Get(req *UserGetParam) (*entity.User, error) {
	user := new(entity.User)
	err := c.repository.Find(req.Email, user)
	if err != nil {
		return nil, err
	}

	res := &entity.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

func (c *userUsecase) GetCurrentUser(ctx *gin.Context, req *UserGetParam) error {
	return c.authHelper.GetAuthUser(ctx, &entity.User{
		Email: req.Email,
	})
}
