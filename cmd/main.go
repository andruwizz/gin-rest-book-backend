package main

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/config"
)

func main() {
	// Init Database
	db := config.NewDatabase()

	// Init App
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:         db,
		RouteGroup: app.Group("/api/v1"),
	})

	app.Run(":3000")
}
