package router

import (
	"notebook-backend/controller"
	"notebook-backend/repository"
	"notebook-backend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	api := r.Group("/api")
	userRoutes := api.Group("/users")
	{
		userRoutes.GET("/", userController.GetAllUsers)
	}
}
