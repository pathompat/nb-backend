package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductionHandler struct {
	service service.ProductionService
}

func NewProductionHandler(service service.ProductionService) *ProductionHandler {
	return &ProductionHandler{service: service}
}

// ProductionHandler GetProductionByID
//
// @id				GetProductionByID
// @tags			production
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			productionId path string false "production ID"
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.ProductionResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/production/{productionId} [GET]
func (c *ProductionHandler) GetProductionByID(ctx *gin.Context) {
	productionIDParam := ctx.Param("productionId")

	productionID, err := strconv.ParseUint(productionIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	production, err := c.service.GetProductionByID(uint(productionID))
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, production)
}

// ProductionHandler UpdateStatusProductionByID
//
// @id				UpdateStatusProductionByID
// @tags			production
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			productionId path string false "production ID"
// @Param			itemId path string false "item ID"
// @Param			UpdateStatusItemProductionDTO body dto.UpdateStatusItemProduction false "Update status item production request"
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.ProductionItemResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/production/{productionId}/item/{itemId} [PUT]
func (c *ProductionHandler) UpdateStatusProductionByID(ctx *gin.Context) {
	var statusInput dto.UpdateStatusItemProduction
	if err := ctx.ShouldBindJSON(&statusInput); err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	productionIDParam := ctx.Param("productionId")
	itemIDParam := ctx.Param("itemId")

	productionID, err := strconv.ParseUint(productionIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	itemID, err := strconv.ParseUint(itemIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	production, err := c.service.UpdateStatusProductionByID(uint(productionID), uint(itemID), statusInput)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, production)
}
