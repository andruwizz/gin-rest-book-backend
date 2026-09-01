package routes

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.RouterGroup, h *handler.BookHandler) {
	books := r.Group("/books")
	books.Use(middleware.Authenticated())
	books.POST("/", h.Create)
	books.GET("/", h.List)
	books.GET("/:id", h.Find)
	books.PUT("/:id", h.Update)
	books.DELETE("/:id", h.Delete)
}
