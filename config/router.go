package config

import (
	"log/slog"
	"net/http"
	"notebook-backend/controller"
	"notebook-backend/repository"
	"notebook-backend/service"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)

	loginService := service.NewLoginService(userRepo)
	loginController := controller.NewLoginController(loginService)

	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	api := r.Group("/api")
	api.POST("/login", loginController.Login)

	api.Use(authMiddleware())
	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:userId", userController.UpdateUser)
		userRoutes.DELETE("/:userId", userController.DeleteUser)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		splitToken := strings.Split(authorization, "Bearer ")
		tokenString := splitToken[1]

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			slog.Error(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": "Unauthorized"})
			c.Abort()
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			slog.Error("token invalid")
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
