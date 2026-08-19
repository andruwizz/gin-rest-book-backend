package main

import (
	"github.com/andruwizz/gin-collective-library-backend/cmd/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Init Route Config
	api := router.Group("/api/v1")
	routes.Setup(api)

	router.Run(":3000")
}
