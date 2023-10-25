package http

import (
	"hub-connect/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

func DefineAPIRoutes(r *gin.Engine, uc controller.UserController, tc controller.TeamController, hc controller.HubController) {
	// Set up API routes using Gin's router groups.
	v1 := r.Group("/v1")

	// private API
	v1.Use(AuthMiddleware)

	hubRoutes := v1.Group("/hubs")
	{
		hubRoutes.POST("", hc.CreateHub)
		hubRoutes.POST("/search", hc.SearchHubs)
		hubRoutes.GET("/:hubID", hc.GetHubByID)
	}

	teamRoutes := v1.Group("/teams")
	{
		teamRoutes.POST("", tc.CreateTeam)
		teamRoutes.POST("/search", tc.SearchTeams)
		teamRoutes.GET("/:teamID", tc.GetTeamByID)
		teamRoutes.GET("/:teamID/hubs/:hubID", tc.TeamJoinHub)
	}

	userRoutes := v1.Group("/users")
	{
		userRoutes.POST("", uc.CreateUser)
		userRoutes.GET("/:userID", uc.GetUserByID)
		userRoutes.GET("/:userID/teams/:teamID", uc.UserJoinTeam)
	}
}
