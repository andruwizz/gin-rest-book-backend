package routes

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func LendingRouter(r *gin.RouterGroup, h *handler.LendingHandler) {
	lendings := r.Group("/lendings")
	lendings.GET("/", h.List)
}
