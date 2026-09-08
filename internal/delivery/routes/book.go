package routes

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/middleware"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	"github.com/gin-gonic/gin"
)

func BookRouter(r *gin.RouterGroup, h *handler.BookHandler, ah *helper.AuthHelper) {
	books := r.Group("/books")
	books.Use(middleware.Authenticated(ah))
	books.POST("/", h.Create)
	books.GET("/", h.List)
	books.GET("/:id", h.Find)
	books.PUT("/:id", h.Update)
	books.DELETE("/:id", h.Delete)
}
