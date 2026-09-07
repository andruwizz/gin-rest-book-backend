package main

import (
	"fmt"

	"github.com/andruwizz/gin-rest-book-backend/internal/config"
)

func main() {
	// Init Environment
	env := config.NewEnv()

	// Init Database
	db := config.NewDatabase(env)

	// Init App
	app := config.NewGin()

	config.Bootstrap(&config.BootstrapConfig{
		Env: env,
		DB:  db,
		App: app,
	})

	appAddress := fmt.Sprintf(":%s", env.AppPort)
	app.Run(appAddress)
}
