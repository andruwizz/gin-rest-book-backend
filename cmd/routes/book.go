package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.RouterGroup, h *handler.BookHandler) {
	books := r.Group("/books")
	books.GET("/", h.List)
}
