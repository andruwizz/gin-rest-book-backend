package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	// Setup Handler
	bookHandler := handler.NewBookHandler()
	userHandler := handler.NewUserHandler()

	// Setup Router
	BookRouter(app, bookHandler)
	UserRouter(app, userHandler)
}
