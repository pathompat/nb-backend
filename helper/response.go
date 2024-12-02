package helper

import (
	"github.com/gin-gonic/gin"
)

type ApiErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ApiSuccessResponse struct {
	ApiErrorResponse
	Data interface{} `json:"data"`
}

func ErrorResponse(ctx *gin.Context, code int, err error) {
	ctx.Error(err)
	ctx.AbortWithStatusJSON(code, ApiErrorResponse{
		Code:    code,
		Message: err.Error(),
	})
}

func SuccessResponse(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, ApiSuccessResponse{
		ApiErrorResponse: ApiErrorResponse{Code: code, Message: "Success"},
		Data:             data,
	})
}

func SuccessResponseWithMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, ApiErrorResponse{
		Code:    code,
		Message: message,
	})
}
