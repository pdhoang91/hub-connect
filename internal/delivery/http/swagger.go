package http

import (
	"hub-connect/config"
	"hub-connect/docs"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureSwagger(cfg *config.Config, r *gin.Engine) {
	// Configure Swagger documentation
	docs.SwaggerInfo.Title = "API Documentations"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = cfg.SWAGGER_DOMAIN
	docs.SwaggerInfo.Schemes = []string{"http"}

	// Enable Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
