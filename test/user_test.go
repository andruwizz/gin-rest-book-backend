package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	usecaseMock "github.com/andruwizz/gin-rest-book-backend/internal/usecase/user/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewUserMockSetup() (*handler.UserHandler, *helper.AuthHelper) {
	authHelper := helper.NewAuthHelper("Book REST Backend", "userInfo", 1, "test-signature-key")
	userService := usecaseMock.NewMockUserUsecase()
	userHandler := handler.NewUserHandler(userService, authHelper)

	return userHandler, authHelper
}

func TestRegisterUser(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	userHandler, authHelper := NewUserMockSetup()
	routes.UserRouter(rg, userHandler, authHelper)

	// Act
	w := httptest.NewRecorder()

	payload := request.UserCreate{Name: "John Doe", Email: "john.doe@example.com", Password: "Password123"}
	reqPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/register", bytes.NewBuffer(reqPayload))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.User]
	err := json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, response.DataResponse[entity.User]{
		Success: true,
		Data: entity.User{
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		},
	}, res)
}

func TestLoginUser(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	userHandler, authHelper := NewUserMockSetup()
	routes.UserRouter(rg, userHandler, authHelper)

	// Act
	w := httptest.NewRecorder()

	payload := request.UserLogin{Email: "john.doe@example.com", Password: "Password123"}
	reqPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/login", bytes.NewBuffer(reqPayload))
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.AuthToken]
	err := json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.DataResponse[entity.AuthToken]{
		Success: true,
		Data: entity.AuthToken{
			Token: "mock.jwt.token",
		},
	}, res)
}

func TestCurrentUser(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	userHandler, authHelper := NewUserMockSetup()
	routes.UserRouter(rg, userHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/current", nil)
	req.Header.Set("Authorization", "Bearer "+token.Token)
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.User]
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.DataResponse[entity.User]{
		Success: true,
		Data: entity.User{
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Name:      "John Doe",
			Email:     "john.doe@example.com",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		},
	}, res)
}

func TestCurrentUser_Unauthorized(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	userHandler, authHelper := NewUserMockSetup()
	routes.UserRouter(rg, userHandler, authHelper)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/current", nil)
	app.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
