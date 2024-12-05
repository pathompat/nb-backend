package handler

import (
	"net/http"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	service service.SchoolService
}

func NewSchoolHandler(service service.SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

// SchoolHandler godoc
//
// @id				GetAllUsers
// @tags			users
// @security	JWTToken
// @accept		json
// @produce		json
//
// @query			userId path string false "User's UUID"
//
// @response 200 {object} helper.ApiSuccessResponse{data=[]dto.SchoolResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/schools [GET]
func (c *SchoolHandler) GetSchoolByUserId(ctx *gin.Context) {
	userID := ctx.Query("userId")

	school, err := c.service.GetSchoolByUserId(userID)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, school)
}
