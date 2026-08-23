package routes

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func InventoryRouter(r *gin.RouterGroup, h *handler.InventoryHandler) {
	inventories := r.Group("/inventories")
	inventories.GET("/", h.List)
}
