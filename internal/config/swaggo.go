package config

import (
	"fmt"

	"github.com/andruwizz/gin-rest-book-backend/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func NewSwaggo(app *gin.Engine, env *Env) {

	// Set Swagger Info
	docs.SwaggerInfo.Title = "Book REST API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a Book REST API"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", env.AppHost, env.SwaggoPort)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Set Swagger Path
	app.GET("/api/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
}
