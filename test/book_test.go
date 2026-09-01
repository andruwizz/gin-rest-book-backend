package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andruwizz/gin-rest-book-backend/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestListBookUnauthorized(t *testing.T) {
	app := config.NewGin()
	config.Bootstrap(&config.BootstrapConfig{
		App: app,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
