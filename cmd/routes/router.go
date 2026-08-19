package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func Setup(rg *gin.RouterGroup) {
	// Setup Handler
	bookHandler := handler.NewBookHandler()
	userHandler := handler.NewUserHandler()

	// Setup Router
	BookRouter(rg, bookHandler)
	UserRouter(rg, userHandler)
}
