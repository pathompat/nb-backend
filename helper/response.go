package helper

import (
	"github.com/gin-gonic/gin"
)

type ApiErrorResponse struct {
	Code    int    `json:"code" example:"400"`            // HTTP status
	Message string `json:"message" example:"Bad request"` // Error message
}

type ApiSuccessResponse struct {
	Code    int         `json:"code" example:"200"`        // HTTP status
	Message string      `json:"message" example:"Success"` // Return message
	Data    interface{} `json:"data"`                      // Returning data
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
		Code:    code,
		Message: "Success",
		Data:    data,
	})
}

func SuccessResponseWithMessage(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, ApiErrorResponse{
		Code:    code,
		Message: message,
	})
}
