package config

import (
	"github.com/andruwizz/gin-book-sharing-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewSwaggo(app *gin.Engine) {

	// Set Swagger Info
	docs.SwaggerInfo.Title = "Book Sharing API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a Book Sharing API"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Set Swagger Path
	app.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
