package config

import (
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/handler"
	"github.com/andruwizz/gin-book-sharing-backend/internal/delivery/routes"
	"github.com/andruwizz/gin-book-sharing-backend/internal/repository"
	"github.com/andruwizz/gin-book-sharing-backend/internal/usecase"
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

	// Setup Service
	bookService := usecase.NewBookUsecase(bookRepository)

	// Setup Handler
	bookHandler := handler.NewBookHandler(bookService)

	// Setup Router
	rg := config.App.Group("/api/v1")
	routes.BookRouter(rg, bookHandler)

	// Setup API Documentation
	NewSwaggo(config.App)
}
