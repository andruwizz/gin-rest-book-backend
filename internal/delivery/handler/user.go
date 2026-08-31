package handler

import (
	"errors"
	"net/http"

	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/request"
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/response"
	"github.com/andruwizz/gin-book-sharing-backend/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

// Register godoc
// @Summary Register a new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param Body body request.UserCreate true "Request body"
// @Success 201 {object} response.DataResponse{data=entity.User}
// @Failure 400 {object} response.ErrorResponse
// @Router /users/register [POST]
func (u *UserHandler) Register(ctx *gin.Context) {
	req := new(request.UserCreate)
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := u.usecase.Create(req)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.UserResource(ctx, http.StatusCreated, res)
}

// Login godoc
// @Summary Login as user
// @Description Login as registered user
// @Tags User
// @Accept json
// @Produce json
// @Param Body body request.UserLogin true "Request body"
// @Success 200 {object} response.DataResponse{data=entity.AuthToken}
// @Failure 400 {object} response.ErrorResponse
// @Router /users/login [POST]
func (u *UserHandler) Login(ctx *gin.Context) {
	req := new(request.UserLogin)
	if err := ctx.ShouldBind(req); err != nil {
		response.ErrorValidation(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := u.usecase.Login(req)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.UserAuthResource(ctx, http.StatusOK, res)
}

// Current godoc
// @Summary Get current user
// @Description Show current authenticated user
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.DataResponse{data=entity.User}
// @Failure 400 {object} response.ErrorResponse
// @Router /users/current [GET]
func (u *UserHandler) Current(ctx *gin.Context) {
	req := new(request.UserGet)
	userInfo, ok := ctx.Get("userInfo")
	if !ok {
		response.ErrorResource(ctx, http.StatusUnauthorized, errors.New("Unauthenticated"))
		return
	}

	userInfoMap := userInfo.(map[string]string)
	req.Email = userInfoMap["email"]

	res, err := u.usecase.Get(req)
	if err != nil {
		response.ErrorResource(ctx, http.StatusBadRequest, err)
		return
	}

	response.UserResource(ctx, http.StatusOK, res)
}
