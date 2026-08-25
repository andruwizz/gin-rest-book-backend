package main

import (
	_ "github.com/andruwizz/gin-book-sharing-backend/docs"
	"github.com/andruwizz/gin-book-sharing-backend/internal/config"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book Sharing API
// @version 1.0
// @description This is a Book Sharing API

// @license.name MIT
// @license.url https://mit-license.org/

// @host localhost:3000
// @BasePath /api/v1
func main() {
	// Init Database
	db := config.NewDatabase()

	// Init App
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:         db,
		RouteGroup: app.Group("/api/v1"),
	})

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.Run(":3000")
}
