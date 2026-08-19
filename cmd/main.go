package main

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// Init Route Config
	routes.Setup(app)

	app.Run(":3000")
}
