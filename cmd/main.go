package main

import (
	"fmt"
	"os"

	"github.com/andruwizz/gin-book-sharing-backend/internal/config"
)

func main() {
	// Init Database
	db := config.NewDatabase()

	// Init App
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		DB:  db,
		App: app,
	})

	appAddress := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))
	app.Run(appAddress)
}
