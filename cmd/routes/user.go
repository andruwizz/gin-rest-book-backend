package routes

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(app *gin.Engine, h *handler.UserHandler) {
	app.GET("/users", h.List)
}
