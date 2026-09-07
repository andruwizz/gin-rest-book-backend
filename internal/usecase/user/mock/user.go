package mock

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/user"
)

type UserUsecase interface {
	Create(req *user.UserCreateParam) (*entity.User, error)
	Login(req *user.UserLoginParam) (*entity.AuthToken, error)
	Get(req *user.UserGetParam) (*entity.User, error)
}

type mockUserUsecase struct{}

func NewMockUserUsecase() UserUsecase {
	return &mockUserUsecase{}
}

func (c *mockUserUsecase) Create(req *user.UserCreateParam) (*entity.User, error) {
	return &entity.User{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}, nil
}

func (c *mockUserUsecase) Login(req *user.UserLoginParam) (*entity.AuthToken, error) {
	return &entity.AuthToken{
		Token: "mock.jwt.token",
	}, nil
}

func (c *mockUserUsecase) Get(req *user.UserGetParam) (*entity.User, error) {
	return &entity.User{
		Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
		Name:      "John Doe",
		Email:     req.Email,
		CreatedAt: 1788410288537,
		UpdatedAt: 1788410288537,
	}, nil
}
