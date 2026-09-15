package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/response"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewDefaultMockSetup() *handler.DefaultHandler {
	defaultHandler := handler.NewDefaultHandler()

	return defaultHandler
}
func TestIndex(t *testing.T) {
	// Arrange
	app := gin.Default()
	defaultHandler := NewDefaultMockSetup()
	routes.DefaultRouter(app, defaultHandler)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	app.ServeHTTP(w, req)

	// Assert
	var res response.EmptyResponse
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, response.EmptyResponse{
		Success: true,
	}, res)
}

func TestNotFound(t *testing.T) {
	// Arrange
	app := gin.Default()
	defaultHandler := NewDefaultMockSetup()
	routes.DefaultRouter(app, defaultHandler)

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/somewhere", nil)
	app.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
}
