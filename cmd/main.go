package main

import (
	"github.com/andruwizz/gin-collective-library-backend/internal/config"
)

func main() {
	// Init Database
	db := config.NewDatabase()

	// Init App
	app := config.NewGin()

	bootstrap := config.BootstrapConfig{
		DB:         db,
		RouteGroup: app.Group("/api/v1"),
	}
	config.Bootstrap(&bootstrap)

	app.Run(":3000")
}
