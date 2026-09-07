package handler

import (
	"net/http"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase    user.UserUsecase
	authHelper *helper.AuthHelper
}

func NewUserHandler(usecase user.UserUsecase, authHelper *helper.AuthHelper) *UserHandler {
	return &UserHandler{usecase, authHelper}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param Body body request.UserCreate true "Request body"
// @Success 201 {object} response.DataResponse[entity.User]
// @Failure 400 {object} response.ErrorResponse
// @Router /users/register [POST]
func (u *UserHandler) Register(ctx *gin.Context) {
	req := new(request.UserCreate)
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &user.UserCreateParam{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.usecase.Create(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusCreated, res)
}

// Login godoc
// @Summary Login as user
// @Description Login as registered user
// @Tags User
// @Accept json
// @Produce json
// @Param Body body request.UserLogin true "Request body"
// @Success 200 {object} response.DataResponse[entity.AuthToken]
// @Failure 400 {object} response.ErrorResponse
// @Router /users/login [POST]
func (u *UserHandler) Login(ctx *gin.Context) {
	req := new(request.UserLogin)
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &user.UserLoginParam{
		Email:    req.Email,
		Password: req.Password,
	}
	res, err := u.usecase.Login(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusOK, res)
}

// Current godoc
// @Summary Get current user
// @Description Show current authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.DataResponse[entity.User]
// @Failure 400 {object} response.ErrorResponse
// @Router /users/current [GET]
func (u *UserHandler) Current(ctx *gin.Context) {
	current := new(entity.User)
	err := u.authHelper.GetAuthUser(ctx, current)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	dto := &user.UserGetParam{
		Name:  current.Name,
		Email: current.Email,
	}
	res, err := u.usecase.Get(dto)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.Resource(ctx, http.StatusOK, res)
}
