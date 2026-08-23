package routes

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, h *handler.UserHandler) {
	users := r.Group("/users")
	users.GET("/", h.List)
}
