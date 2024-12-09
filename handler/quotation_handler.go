package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"
	"strconv"

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

// QuotationHandler GetQuotationByID
//
// @id				GetQuotationByID
// @tags			quotations
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			quotationId path string false "quotation ID"
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.QuotationResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
// @router			/quotation/{quotationId} [GET]
func (c *QuotationHandler) GetQuotationByID(ctx *gin.Context) {
	quotationIDParam := ctx.Param("quotationId")

	quotationID, err := strconv.ParseUint(quotationIDParam, 10, 32)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	quotations, err := c.service.GetQuotationByID(uint(quotationID))
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

// QuotationHandler UpdateQuotation
//
// @id				UpdateQuotation
// @tags			quotations
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			quotationId path int false "Quotation id"
// @Param			updateQuotationDTO body dto.UpdateQuotation false "Update quotation request"
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.QuotationResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/quotation/{quotationId} [PUT]
func (c *QuotationHandler) UpdateQuotation(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		helper.ErrorResponse(ctx, http.StatusBadRequest, helper.ErrInvalidPathParam)
		return
	}

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	var request dto.UpdateQuotation
	if err := ctx.ShouldBindJSON(&request); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.service.UpdateQuotation(uint(idInt), request)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, user)
}
