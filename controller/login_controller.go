package controller

import (
	"net/http"
	"notebook-backend/controller/dto"
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
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	token, err := c.service.Login(credential)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
