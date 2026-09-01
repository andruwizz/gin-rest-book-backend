package routes

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, h *handler.UserHandler) {
	users := r.Group("/users")
	users.POST("/register", h.Register)
	users.POST("/login", h.Login)
	users.GET("/current", middleware.Authenticated(), h.Current)
}
