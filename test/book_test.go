package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	repoMock "github.com/andruwizz/gin-rest-book-backend/internal/repository/mock"
	usecaseMock "github.com/andruwizz/gin-rest-book-backend/internal/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewMockSetup() *handler.BookHandler {
	bookRepository := repoMock.NewMockBookRepository()
	bookService := usecaseMock.NewMockBookUsecase(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	return bookHandler
}

// func TestListBookUnauthorized(t *testing.T) {
// 	// Arrange
// 	app := gin.Default()
// 	rg := app.Group("/api/v1")
// 	routes.BookRouter(rg, NewMockSetup())

// 	// Act
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
// 	app.ServeHTTP(w, req)

// 	// Assert
// 	res := `{"success":false,"error":{"message":"authentication token not found"}}`
// 	assert.Equal(t, http.StatusUnauthorized, w.Code)
// 	assert.Equal(t, res, w.Body.String())
// }

func TestListBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	routes.BookRouter(rg, NewMockSetup())

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	app.ServeHTTP(w, req)

	// Assert
	res := `{"success":true,"data":[{"id":"305c8059-28d7-49f7-a15c-acba512f2b0a","title":"Book One","author":"Book Author","created_at":1788410288537,"updated_at":1788410288537},{"id":"348c6370-801d-4fab-80a3-5b3ecbc88760","title":"Book Two","author":"Book Author","created_at":1788410288537,"updated_at":1788410288537},{"id":"5a5a9021-7b19-47f4-bda8-05d2551f8ed8","title":"Book Three","author":"Book Author","created_at":1788410288537,"updated_at":1788410288537}],"meta":{"page":1,"per_page":10,"total":100,"total_pages":10}}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}

func TestFindBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	routes.BookRouter(rg, NewMockSetup())

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/348c6370-801d-4fab-80a3-5b3ecbc88760", nil)
	app.ServeHTTP(w, req)

	// Assert
	res := `{"success":true,"data":{"id":"348c6370-801d-4fab-80a3-5b3ecbc88760","title":"Book Two","author":"Book Author","created_at":1788410288537,"updated_at":1788410288537}}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}

func TestDeleteBook(t *testing.T) {
	// Arrange
	app := gin.Default()
	rg := app.Group("/api/v1")
	routes.BookRouter(rg, NewMockSetup())

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/books/348c6370-801d-4fab-80a3-5b3ecbc88760", nil)
	app.ServeHTTP(w, req)

	// Assert
	res := `{"success":true}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, res, w.Body.String())
}
