package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
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
// @id				GetSchoolByUserId
// @tags			schools
// @security	JWTToken
// @accept		json
// @produce		json
//
// @Param			userId query string false "User's UUID"
//
// @response 200 {object} helper.ApiSuccessResponse{data=[]dto.SchoolResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/school [GET]
func (c *SchoolHandler) GetSchoolByUserId(ctx *gin.Context) {
	userID := ctx.Query("userId")

	school, err := c.service.GetSchoolByUserId(userID)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, school)
}

// SchoolHandler godoc
//
// @id				CreateSchool
// @tags			schools
// @security	JWTToken
// @accept		json
// @produce		json
//
// @Param			createSchoolDTO body dto.CreateSchool false "Create school request"
//
// @response 201 {object} helper.ApiSuccessResponse{data=dto.SchoolResponse} "Created"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/school [POST]
func (c *SchoolHandler) CreateSchool(ctx *gin.Context) {
	var schoolInput dto.CreateSchool
	if err := ctx.ShouldBindJSON(&schoolInput); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	school, err := c.service.CreateSchool(schoolInput)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, school)
}
