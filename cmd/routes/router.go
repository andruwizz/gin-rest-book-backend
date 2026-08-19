package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func Setup(rg *gin.RouterGroup) {
	// Setup Handler
	bookHandler := handler.NewBookHandler()
	inventoryHandler := handler.NewInventoryHandler()
	lendingHandler := handler.NewLendingHandler()
	userHandler := handler.NewUserHandler()

	// Setup Router
	BookRouter(rg, bookHandler)
	InventoryRouter(rg, inventoryHandler)
	LendingRouter(rg, lendingHandler)
	UserRouter(rg, userHandler)
}
