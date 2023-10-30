package controller

import (
	"encoding/json"
	"fmt"
	"hub-connect/internal/delivery/http/controller/mocks"
	"hub-connect/internal/delivery/http/model"
	"hub-connect/internal/entities"
	"hub-connect/internal/usecase"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestTeamController_CreateTeam(t *testing.T) {

	type fields struct {
		teamUseCase usecase.TeamUseCase
	}

	type args struct {
		mockCtx func(ctx *gin.Context)
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		expectedResponse string
		expectedStatus   int
	}{
		// TODO: Add test cases.
		{
			name: "happy case",
			fields: fields{
				teamUseCase: func() usecase.TeamUseCase {
					mockService := &mocks.TeamUseCase{}
					mockService.On("CreateTeam", "Team 1", "Type A").Return(&entities.Team{
						Name: "Team 1",
						Type: "Type A",
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					payload := model.CreateTeamRequest{
						Name: "Team 1",
						Type: "Type A",
					}
					bodyString, _ := json.Marshal(payload)
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/v1/teams",
						strings.NewReader(string(bodyString)),
					)
				},
			},
			expectedResponse: "{\"status\":\"success\",\"code\":200,\"data\":{\"team_id\":0,\"name\":\"Team 1\",\"type\":\"Type A\",\"hub_id\":null,\"created_at\":null,\"updated_at\":null,\"users\":null}}",
			expectedStatus:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c := &TeamController{
				TeamUseCase: tt.fields.teamUseCase,
			}
			tt.args.mockCtx(ctx)
			c.CreateTeam(ctx)
			fmt.Println("tt.expectedResponse", tt.expectedResponse)
			fmt.Println("w.Body.String()", w.Body.String())
			assert.Equal(t, tt.expectedResponse, w.Body.String())
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
