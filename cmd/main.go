package main

import "github.com/andruwizz/gin-book-sharing-backend/internal/config"

func main() {
	// Init Database
	db := config.NewDatabase()

	// Init App
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
	})

	app.Run(":3000")
}
