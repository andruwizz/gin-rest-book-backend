package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestListBookUnauthorized(t *testing.T) {
	// Arrange
	app := config.NewGin()
	config.Bootstrap(&config.BootstrapConfig{
		App: app,
	})

	// Act
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	app.ServeHTTP(w, req)

	// Assert
	res := `{"success":false,"error":{"message":"authentication token not found"}}`
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, res, w.Body.String())
}
