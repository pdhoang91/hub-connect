// Package api defines the handlers for managing hubs through the API endpoints.
package controller

import (
	"hub-connect/internal/delivery/http/model"
	"hub-connect/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HubHandler is responsible for handling HTTP requests related to hubs.
type HubController struct {
	Controller
	HubUseCase usecase.HubUseCase
}

// NewHubController creates a new instance of HubHandler with the provided HubUseCase.
func NewHubController(hubUseCase usecase.HubUseCase) HubController {
	return HubController{HubUseCase: hubUseCase}
}

// CreateHub handles the creation of a new hub through an HTTP POST request.
func (hc *HubController) CreateHub(ctx *gin.Context) {
	var request model.CreateHubRequest
	if hc.WrapBindAndValidate(ctx, &request) {
		hubID, err := hc.HubUseCase.CreateHub(request.Name, request.Location)
		if err != nil {
			hc.Failure(ctx, err)
			return
		}
		hc.Success(ctx, hubID)
	}
}

// GetHubByID retrieves a hub's information by its ID through an HTTP GET request.
func (hc *HubController) GetHubByID(ctx *gin.Context) {
	hubID, err := strconv.Atoi(ctx.Param("hubID"))
	if err != nil {
		//responc(ctx, http.StatusBadRequest, nil, "Invalid hub ID")
		hc.Failure(ctx, err)
		return
	}

	hub, err := hc.HubUseCase.GetHubByID(int(hubID))
	if err != nil {
		//responc(ctx, http.StatusNotFound, nil, "Hub not found")
		hc.Failure(ctx, err)
		return
	}

	//responc(ctx, http.StatusOK, hub, "")
	hc.Success(ctx, hub)
}

// SearchHubs handles searching for hubs based on a keyword through an HTTP POST request.
func (hc *HubController) SearchHubs(ctx *gin.Context) {
	var request model.SearchRequest
	if hc.WrapBindAndValidate(ctx, &request) {
		hubs, err := hc.HubUseCase.SearchHubs(request.Keyword)
		if err != nil {
			//responc(c, http.StatusInternalServerError, nil, "Failed to search hubs")
			hc.Failure(ctx, err)
			return
		}

		//responc(c, http.StatusOK, hubs, "")
		hc.Success(ctx, hubs)
	}
}
