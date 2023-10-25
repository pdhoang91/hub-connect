// Package api defines the handlers for managing users through the API endpoints.
package controller

import (
	"fmt"
	"hub-connect/internal/delivery/http/model"
	"hub-connect/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserController is responsible for handling HTTP requests related to users.
type UserController struct {
	Controller
	UserUseCase usecase.UserUseCase
}

// NewUserController creates a new instance of UserHandler with the provided UserUseCase.
func NewUserController(userUseCase usecase.UserUseCase) *UserController {
	return &UserController{UserUseCase: userUseCase}
}

// CreateUser handles the creation of a new user through an HTTP POST request.
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var request model.CreateUserRequest
	if uc.WrapBindAndValidate(ctx, &request) {
		// Call the UserUseCase to create a new user.
		id, err := uc.UserUseCase.CreateUser(request.Name, request.Email)
		if err != nil {
			uc.Failure(ctx, err)
			return
		}
		uc.Success(ctx, id)
	}
}

// GetUserByID retrieves a user's information by its ID through an HTTP GET request.
func (uc *UserController) GetUserByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		uc.Failure(ctx, err)
		return
	}
	// Call the UserUseCase to get a user by its ID.
	user, err := uc.UserUseCase.GetUserByID(int(userID))
	if err != nil {
		uc.Failure(ctx, err)
		return
	}
	uc.Success(ctx, user)
}

// UserJoinTeam handles associating a user with a team through an HTTP GET request.
func (uc *UserController) UserJoinTeam(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		uc.Failure(ctx, err)
		return
	}
	teamID, err := strconv.Atoi(ctx.Param("teamID"))
	if err != nil {
		uc.Failure(ctx, err)
		return
	}
	// Call the UserUseCase to associate the user with the team.
	err = uc.UserUseCase.UserJoinTeam(int(userID), int(teamID))
	if err != nil {
		uc.Failure(ctx, err)
		return
	}
	uc.Success(ctx, fmt.Sprintf("User [%d] join Team [%d] successfully", userID, teamID))
}
