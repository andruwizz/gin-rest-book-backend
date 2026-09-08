package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/request"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-rest-book-backend/internal/entity"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	usecaseMock "github.com/andruwizz/gin-rest-book-backend/internal/usecase/book/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewBookMockSetup() (*handler.BookHandler, *helper.AuthHelper) {
	authHelper := helper.NewAuthHelper("Book REST Backend", "userInfo", 1, "test-signature-key")
	bookService := usecaseMock.NewMockBookUsecase()
	bookHandler := handler.NewBookHandler(bookService)

	return bookHandler, authHelper
}

func TestCreateBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	bookHandler, authHelper := NewBookMockSetup()
	routes.BookRouter(rg, bookHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()

	payload := request.BookCreate{Title: "Laskar Pelari", Author: "Andria Henrietta"}
	reqPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/books/", bytes.NewBuffer(reqPayload))
	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.Book]
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, response.DataResponse[entity.Book]{
		Success: true,
		Data: entity.Book{
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Title:     "Laskar Pelari",
			Author:    "Andria Henrietta",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		},
	}, res)
}

func TestListBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	bookHandler, authHelper := NewBookMockSetup()
	routes.BookRouter(rg, bookHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	req.Header.Set("Authorization", "Bearer "+token.Token)
	app.ServeHTTP(w, req)

	// Assert
	var res response.ListResponse[entity.Book]
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.ListResponse[entity.Book]{
		Success: true,
		Data: []entity.Book{
			{
				Id:        "305c8059-28d7-49f7-a15c-acba512f2b0a",
				Title:     "Book One",
				Author:    "Book Author",
				CreatedAt: 1788410288537,
				UpdatedAt: 1788410288537,
			},
			{
				Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
				Title:     "Book Two",
				Author:    "Book Author",
				CreatedAt: 1788410288537,
				UpdatedAt: 1788410288537,
			},
			{
				Id:        "5a5a9021-7b19-47f4-bda8-05d2551f8ed8",
				Title:     "Book Three",
				Author:    "Book Author",
				CreatedAt: 1788410288537,
				UpdatedAt: 1788410288537,
			},
		},
		Meta: response.ListMeta{
			Page:       1,
			PerPage:    10,
			Total:      100,
			TotalPages: 10,
		},
	}, res)
}

func TestFindBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	bookHandler, authHelper := NewBookMockSetup()
	routes.BookRouter(rg, bookHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/348c6370-801d-4fab-80a3-5b3ecbc88760", nil)
	req.Header.Set("Authorization", "Bearer "+token.Token)
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.Book]
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.DataResponse[entity.Book]{
		Success: true,
		Data: entity.Book{
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Title:     "Book Two",
			Author:    "Book Author",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		},
	}, res)
}

func TestUpdateBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	bookHandler, authHelper := NewBookMockSetup()
	routes.BookRouter(rg, bookHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()

	payload := request.BookCreate{Title: "Laskar Pelari", Author: "Andria Hirata"}
	reqPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/books/348c6370-801d-4fab-80a3-5b3ecbc88760", bytes.NewBuffer(reqPayload))
	req.Header.Set("Authorization", "Bearer "+token.Token)
	req.Header.Set("Content-Type", "application/json")
	app.ServeHTTP(w, req)

	// Assert
	var res response.DataResponse[entity.Book]
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.DataResponse[entity.Book]{
		Success: true,
		Data: entity.Book{
			Id:        "348c6370-801d-4fab-80a3-5b3ecbc88760",
			Title:     "Laskar Pelari",
			Author:    "Andria Hirata",
			CreatedAt: 1788410288537,
			UpdatedAt: 1788410288537,
		},
	}, res)
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	bookHandler, authHelper := NewBookMockSetup()
	routes.BookRouter(rg, bookHandler, authHelper)

	token, err := authHelper.EncodeAuthToken(&entity.User{Name: "John Doe", Email: "john.doe@example.com"})
	assert.NoError(t, err)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/books/348c6370-801d-4fab-80a3-5b3ecbc88760", nil)
	req.Header.Set("Authorization", "Bearer "+token.Token)
	app.ServeHTTP(w, req)

	// Assert
	var res response.EmptyResponse
	err = json.Unmarshal(w.Body.Bytes(), &res)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.EmptyResponse{
		Success: true,
	}, res)
}
