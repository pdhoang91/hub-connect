// Package api defines the handlers for managing teams through the API endpoints.
package controller

import (
	"fmt"
	"hub-connect/internal/delivery/http/model"
	"hub-connect/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TeamHandler is responsible for handling HTTP requests related to teams.
type TeamController struct {
	Controller
	TeamUseCase usecase.TeamUseCase
}

// NewTeamController creates a new instance of TeamHandler with the provided TeamUseCase.
func NewTeamController(teamUseCase usecase.TeamUseCase) TeamController {
	return TeamController{TeamUseCase: teamUseCase}
}

// CreateTeam handles the creation of a new team through an HTTP POST request.
func (tc *TeamController) CreateTeam(ctx *gin.Context) {
	var request model.CreateTeamRequest
	if tc.WrapBindAndValidate(ctx, &request) {
		// Call the TeamUseCase to create a new team.
		id, err := tc.TeamUseCase.CreateTeam(request.Name, request.Type)
		if err != nil {
			tc.Failure(ctx, err)
			return
		}
		tc.Success(ctx, id)
	}
}

// GetTeamByID retrieves a team's information by its ID through an HTTP GET request.
func (tc *TeamController) GetTeamByID(ctx *gin.Context) {
	teamID, err := strconv.Atoi(ctx.Param("teamID"))
	if err != nil {
		tc.Failure(ctx, err)
		return
	}
	// Call the TeamUseCase to get a team by its ID.
	team, err := tc.TeamUseCase.GetTeamByID(int(teamID))
	if err != nil {
		tc.Failure(ctx, err)
		return
	}
	tc.Success(ctx, team)
}

// TeamJoinHub handles associating a team with a hub through an HTTP GET request.
func (tc *TeamController) TeamJoinHub(ctx *gin.Context) {
	teamID, err := strconv.Atoi(ctx.Param("teamID"))
	if err != nil {
		tc.Failure(ctx, err)
		return
	}
	hubID, err := strconv.Atoi(ctx.Param("hubID"))
	if err != nil {
		tc.Failure(ctx, err)
		return
	}
	// Call the TeamUseCase to associate the team with the hub.
	err = tc.TeamUseCase.TeamJoinHub(int(teamID), int(hubID))
	if err != nil {
		tc.Failure(ctx, err)
		return
	}
	tc.Success(ctx, fmt.Sprintf("Team [%d] join Hub [%d] successfully", teamID, hubID))
}

// SearchTeams handles searching for teams based on a keyword through an HTTP POST request.
func (tc *TeamController) SearchTeams(ctx *gin.Context) {
	var request model.SearchRequest
	if tc.WrapBindAndValidate(ctx, &request) {
		// Call the TeamUseCase to search for teams based on the provided keyword.
		teams, err := tc.TeamUseCase.SearchTeams(request.Keyword)
		if err != nil {
			tc.Failure(ctx, err)
			return
		}
		tc.Success(ctx, teams)
	}
}
