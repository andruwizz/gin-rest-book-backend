package routes

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func DefaultRouter(app *gin.Engine, h *handler.DefaultHandler) {
	app.GET("/", h.Index)
	app.NoRoute(h.NotFound)
}
