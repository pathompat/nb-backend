package handler

import (
	"net/http"
	"notebook-backend/handler/dto"
	"notebook-backend/helper"
	"notebook-backend/service"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// UserHandler GetAllUsers
//
// @id				GetAllUsers
// @tags			users
// @security	JwtToken
// @accept		json
// @produce		json
//
// @response 200 {object} helper.ApiSuccessResponse{data=[]dto.UserResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
//	@router			/user [GET]
func (c *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers()
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, users)
}

// UserHandler godoc
//
// @id				GetInfoUser
// @tags			users
// @security	JwtToken
// @accept		json
// @produce		json
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.UserResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
//
//	@router			/user/info [GET]
func (c *UserHandler) GetInfoUser(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	claimsMap := claims.(jwt.MapClaims)
	userId, ok := claimsMap["userId"].(string)
	if !ok {
		helper.ErrorResponse(ctx, http.StatusBadRequest, helper.ErrInvalidToken)
		return
	}
	users, err := c.service.GetInfoUser(userId)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	helper.SuccessResponse(ctx, http.StatusOK, users)
}

// UserHandler godoc
//
// @id				CreateUser
// @tags			users
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			createUserDTO body dto.CreateUser false "Create user request"
//
// @response 201 {object} helper.ApiSuccessResponse{data=dto.UserResponse} "Created"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/user [POST]
func (c *UserHandler) CreateUser(ctx *gin.Context) {
	var userInput dto.CreateUser
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

// UserHandler UpdateUser
//
// @id				UpdateUser
// @tags			users
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			updateUserDTO body dto.UpdateUser false "Update user request"
//
// @response 200 {object} helper.ApiSuccessResponse{data=dto.UserResponse} "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/user/{userId} [PUT]
func (c *UserHandler) UpdateUser(ctx *gin.Context) {
	var userInput dto.UpdateUser
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

// UserHandler DeleteUser
//
// @id				DeleteUser
// @tags			users
// @security	JwtToken
// @accept		json
// @produce		json
//
// @Param			userId path string false "User's UUID"
//
// @response 200 "OK"
// @response 400 "Bad request"
// @response 401 "Unauthorized"
// @response 500 "Internal Server Error"
//
//	@router			/user/{userId} [DELETE]
func (c *UserHandler) DeleteUser(ctx *gin.Context) {

	userID := ctx.Param("userId")

	err := c.service.DeleteUser(userID)
	if err != nil {
		helper.ErrorResponse(ctx, http.StatusUnauthorized, err)
		return
	}

	helper.SuccessResponseWithMessage(ctx, http.StatusOK, "Delete user success")
}
