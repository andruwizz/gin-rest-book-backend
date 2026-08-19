package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func LendingRouter(r *gin.RouterGroup, h *handler.LendingHandler) {
	lendings := r.Group("/lendings")
	lendings.GET("/", h.List)
}
