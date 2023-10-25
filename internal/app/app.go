// Package app initializes and configures the application, setting up routes, handlers, and database connections.
package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"hub-connect/config"
	"hub-connect/internal/delivery/http"
	"hub-connect/internal/delivery/http/controller"
	repo "hub-connect/internal/repository"
	uc "hub-connect/internal/usecase"
)

func InitializeHTTPServer(cfg *config.Config, db *gorm.DB) *gin.Engine {
	// Create a new Gin router
	r := gin.New()

	// Enable CORS with custom settings
	http.ConfigureCORS(r)

	// Configure Swagger documentation
	http.ConfigureSwagger(cfg, r)

	// Create repositories and use cases.
	hubRepo := repo.NewHubRepository(db)
	teamRepo := repo.NewTeamRepository(db)
	userRepo := repo.NewUserRepository(db)

	hubUseCase := uc.NewHubUseCase(hubRepo)
	teamUseCase := uc.NewTeamUseCase(teamRepo, hubRepo)
	userUseCase := uc.NewUserUseCase(userRepo, teamRepo)

	// Create API handlers with the associated use cases.
	hubController := controller.NewHubController(hubUseCase)
	teamController := controller.NewTeamController(teamUseCase)
	userController := controller.NewUserController(userUseCase)

	http.DefineAPIRoutes(r, *userController, teamController, hubController)

	return r
}
