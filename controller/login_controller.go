package controller

import (
	"net/http"
	"notebook-backend/controller/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service service.LoginService
}

func NewLoginController(service service.LoginService) *LoginController {
	return &LoginController{service: service}
}

func (c *LoginController) Login(ctx *gin.Context) {
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
