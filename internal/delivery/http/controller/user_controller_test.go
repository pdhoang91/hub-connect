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

func TestUserController_CreateUser(t *testing.T) {

	type fields struct {
		userUseCase usecase.UserUseCase
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
				userUseCase: func() usecase.UserUseCase {
					mockService := &mocks.UserUseCase{}
					mockService.On("CreateUser", "User 1", "email@gmail.com").Return(&entities.User{
						Name:  "User 1",
						Email: "email@gmail.com",
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					payload := model.CreateUserRequest{
						Name:  "User 1",
						Email: "email@gmail.com",
					}
					bodyString, _ := json.Marshal(payload)
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/v1/users",
						strings.NewReader(string(bodyString)),
					)
				},
			},
			expectedResponse: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":0,\"name\":\"User 1\",\"email\":\"email@gmail.com\",\"team_id\":null,\"created_at\":null,\"updated_at\":null}}",
			expectedStatus:   http.StatusOK,
		},
		{
			name: "case invalid payload",
			fields: fields{
				userUseCase: func() usecase.UserUseCase {
					mockService := &mocks.UserUseCase{}
					mockService.On("CreateUser", "User 1", "email@gmail.com").Return(&entities.User{
						Name:  "User 1",
						Email: "email@gmail.com",
					}, nil)
					return mockService
				}(),
			},
			args: args{
				mockCtx: func(ctx *gin.Context) {
					ctx.Request = httptest.NewRequest(
						http.MethodPost,
						"/v1/users",
						strings.NewReader(""),
					)
				},
			},
			expectedResponse: "{\"status\":\"error\",\"code\":400,\"message\":\"unexpected end of JSON input\"}",
			expectedStatus:   http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			c := &UserController{
				UserUseCase: tt.fields.userUseCase,
			}
			tt.args.mockCtx(ctx)
			c.CreateUser(ctx)
			fmt.Println("tt.expectedResponse", tt.expectedResponse)
			fmt.Println("w.Body.String()", w.Body.String())
			assert.Equal(t, tt.expectedResponse, w.Body.String())
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
