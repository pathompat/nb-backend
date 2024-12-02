package controller

import (
	"net/http"
	"notebook-backend/controller/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userInput dto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	user, err := c.service.CreateUser(userInput)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusCreated, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var userInput dto.UpdateUserDTO
	if err := ctx.ShouldBindJSON(&userInput); err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	userID := ctx.Param("userId")

	user, err := c.service.UpdateUser(userID, userInput)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {

	userID := ctx.Param("userId")

	err := c.service.DeleteUser(userID)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponseWithMessage(ctx, http.StatusOK, "Delete user success")
}
