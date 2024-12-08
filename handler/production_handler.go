package handler

import (
	"net/http"
	// "notebook-backend/handler/dto"
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
