package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func BookRouter(app *gin.Engine, h *handler.BookHandler) {
	app.GET("/books", h.List)
}
