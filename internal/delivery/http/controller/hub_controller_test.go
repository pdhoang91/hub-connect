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

func TestHubController_CreateHub(t *testing.T) {

	type fields struct {
		hubUseCase usecase.HubUseCase
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
				hubUseCase: func() usecase.HubUseCase {
					mockService := &mocks.HubUseCase{}
					mockService.On("CreateHub", "Hub 1", "Location A").Return(&entities.Hub{
						Name:     "Hub 1",
						Location: "Location A",
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					payload := model.CreateHubRequest{
						Name:     "Hub 1",
						Location: "Location A",
					}
					bodyString, _ := json.Marshal(payload)
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/v1/hubs",
						strings.NewReader(string(bodyString)),
					)
				},
			},
			expectedResponse: "{\"status\":\"success\",\"code\":200,\"data\":{\"hub_id\":0,\"name\":\"Hub 1\",\"location\":\"Location A\",\"created_at\":null,\"updated_at\":null,\"teams\":null}}",
			expectedStatus:   http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				w := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(w)

				c := &HubController{
					HubUseCase: tt.fields.hubUseCase,
				}
				tt.args.mockCtx(ctx)
				c.CreateHub(ctx)
				fmt.Println("tt.expectedResponse", tt.expectedResponse)
				fmt.Println("w.Body.String()", w.Body.String())
				assert.Equal(t, tt.expectedResponse, w.Body.String())
				assert.Equal(t, tt.expectedStatus, w.Code)
			})
		})
	}
}
