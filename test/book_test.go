package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository/mock"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewMockSetup() *handler.BookHandler {
	bookRepository := mock.NewMockBookRepository()
	bookService := usecase.NewBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	return bookHandler
}

func TestListBookUnauthorized(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	routes.BookRouter(rg, NewMockSetup())

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	app.ServeHTTP(w, req)

	// Assert
	res := `{"success":false,"error":{"message":"authentication token not found"}}`
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, res, w.Body.String())
}
