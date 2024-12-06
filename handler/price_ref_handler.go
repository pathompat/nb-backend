package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type PriceRefHandler struct {
	service service.PriceRefService
}

func NewPriceRefHandler(service service.PriceRefService) *PriceRefHandler {
	return &PriceRefHandler{service: service}
}

// PriceRefHandler godoc
//
// @id				GetPriceRefByUserID
// @tags			priceRef
// @security	JWTToken
// @accept		json
// @produce		json
//
// @Param			userId query string false "User's UUID"
//
// @response 200 {object} helper.ApiSuccessResponse{data=[]dto.PriceRefResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/priceRef [GET]
func (c *PriceRefHandler) GetPriceRefByUserID(ctx *gin.Context) {
	userID := ctx.Query("userId")

	school, err := c.service.GetPriceRefByUserID(userID)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, school)
}

// PriceRefHandler godoc
//
// @id				CreatePriceRef
// @tags			priceRef
// @security	JWTToken
// @accept		json
// @produce		json
//
// @Param			createPriceRefDTO body []dto.CreatePriceRef false "Create priceRef request"
//
// @response 201 {object} helper.ApiSuccessResponse{data=[]dto.PriceRefResponse} "Created"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/priceRef [POST]
func (c *PriceRefHandler) CreatePriceRef(ctx *gin.Context) {
	var priceRefInput []dto.CreatePriceRef
	if err := ctx.ShouldBindJSON(&priceRefInput); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	priceRef, err := c.service.CreatePriceRef(priceRefInput)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, priceRef)
}
