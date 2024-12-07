package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	service service.LoginService
}

func NewLoginHandler(service service.LoginService) *LoginHandler {
	return &LoginHandler{service: service}
}

// LoginHandler Login
//
// @id				Login
// @tags			login
// @security	JwtToken
// @accept		json
// @produce		json
//
// @param			loginRequest body dto.Login true "Login request"
//
// @response 200 {object} dto.ResponseWithToken "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
//	@router			/login [POST]
func (c *LoginHandler) Login(ctx *gin.Context) {
	var credential dto.Login
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	token, err := c.service.Login(credential)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
