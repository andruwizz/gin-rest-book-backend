package config

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-rest-book-backend/internal/helper"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/book"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Env *Env
	DB  *gorm.DB
	App *gin.Engine
}

func Bootstrap(config *BootstrapConfig) {
	// Setup Repository
	bookRepository := repository.NewBookRepository(config.DB)
	userRepository := repository.NewUserRepository(config.DB)

	// Setup helper
	authHelper := helper.NewAuthHelper(config.Env.AppName, config.Env.AuthContextKey, config.Env.AuthTokenDuration, config.Env.AuthSignatureKey)

	// Setup Service
	bookService := book.NewBookUsecase(bookRepository)
	userService := user.NewUserUsecase(userRepository, authHelper)

	// Setup Handler
	bookHandler := handler.NewBookHandler(bookService)
	userHandler := handler.NewUserHandler(userService, authHelper)

	// Setup Router
	rg := config.App.Group("/api/v1")
	routes.BookRouter(rg, bookHandler, authHelper)
	routes.UserRouter(rg, userHandler, authHelper)

	// Setup API Documentation
	NewSwaggo(config.App, config.Env)
}
