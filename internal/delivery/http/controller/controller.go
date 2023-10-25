package controller

import (
	"encoding/json"
	"hub-connect/internal/delivery/http/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
}

func (c *Controller) Success(ctx *gin.Context, data interface{}) {
	response := model.SuccessResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   data,
	}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}

func (c *Controller) Failure(ctx *gin.Context, err error) {
	response := model.ErrorResponse{
		Status:  "error",
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
	ctx.JSON(http.StatusBadRequest, response)
	ctx.Abort()
}

func (c *Controller) WrapBindAndValidate(ctx *gin.Context, request interface{}) bool {
	body, _ := ctx.GetRawData()

	err := json.Unmarshal(body, &request)
	if err != nil {
		c.Failure(ctx, err)
		return false
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		c.Failure(ctx, err)
		return false
	}

	return true
}
