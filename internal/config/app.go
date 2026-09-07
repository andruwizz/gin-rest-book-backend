package config

import (
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-rest-book-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-rest-book-backend/internal/repository"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/book"
	"github.com/andruwizz/gin-rest-book-backend/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB  *gorm.DB
	App *gin.Engine
}

func Bootstrap(config *BootstrapConfig) {
	// Setup Repository
	bookRepository := repository.NewBookRepository(config.DB)
	userRepository := repository.NewUserRepository(config.DB)

	// Setup Service
	bookService := book.NewBookUsecase(bookRepository)
	userService := user.NewUserUsecase(userRepository)

	// Setup Handler
	bookHandler := handler.NewBookHandler(bookService)
	userHandler := handler.NewUserHandler(userService)

	// Setup Router
	rg := config.App.Group("/api/v1")
	routes.BookRouter(rg, bookHandler)
	routes.UserRouter(rg, userHandler)

	// Setup API Documentation
	NewSwaggo(config.App)
}
