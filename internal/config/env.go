package config

import (
	"os"
	"strconv"
)

type Env struct {
	AppName string
	AppHost string
	AppPort string

	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string

	SwaggoPort string

	AuthContextKey    string
	AuthTokenDuration int // Duration in hours
	AuthSignatureKey  string
}

func NewEnv() *Env {
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	duration, _ := strconv.Atoi(os.Getenv("AUTH_TOKEN_DURATION"))

	return &Env{
		AppName: os.Getenv("APP_NAME"),
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),

		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     port,
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),

		SwaggoPort: os.Getenv("SWAGGO_PORT"),

		AuthContextKey:    os.Getenv("AUTH_CONTEXT_KEY"),
		AuthTokenDuration: duration,
		AuthSignatureKey:  os.Getenv("AUTH_SIGNATURE_KEY"),
	}
}
