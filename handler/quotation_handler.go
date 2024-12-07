package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type QuotationHandler struct {
	service service.QuotationService
}

func NewQuotationHandler(service service.QuotationService) *QuotationHandler {
	return &QuotationHandler{service: service}
}

// QuotationHandler GetAllQuotation
//
// @id				GetAllQuotation
// @tags			quotations
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			filter query dto.QuotationFilter false "Filter params"
//
// @response 200 {object} helper.ApiSuccessResponse{data=[]dto.QuotationResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/quotation [GET]
func (c *QuotationHandler) GetAllQuotation(ctx *gin.Context) {
	var filter dto.QuotationFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	quotations, err := c.service.GetAllQuotation(filter)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, quotations)
}

// QuotationHandler CreateQuotation
//
// @id				CreateQuotation
// @tags			quotations
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			createQuotationDTO body dto.CreateQuotation false "Request"
//
// @response 201 {object} helper.ApiSuccessResponse{data=dto.QuotationResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/quotation [POST]
func (c *QuotationHandler) CreateQuotation(ctx *gin.Context) {
	var request dto.CreateQuotation
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	quotation, err := c.service.CreateQuotation(request)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, quotation)
}
