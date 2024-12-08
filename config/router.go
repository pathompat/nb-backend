package config

import (
	"net/http"
	"notebook-backend/handler"
	"notebook-backend/helper"
	"notebook-backend/repository"
	"notebook-backend/service"
	"os"
	"strings"

	_ "notebook-backend/docs" // Swagger generated files

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // Swagger files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	schoolRepo := repository.NewSchoolRepository(db)
	quotationRepo := repository.NewQuotationRepository(db)
	priceRefRepo := repository.NewPriceRefRepository(db)
	productionRepo := repository.NewProductionRepository(db)

	loginService := service.NewLoginService(userRepo)
	loginHandler := handler.NewLoginHandler(loginService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	schoolService := service.NewSchoolService(schoolRepo, userRepo)
	schoolHandler := handler.NewSchoolHandler(schoolService)

	quotationService := service.NewQuotationService(quotationRepo, userRepo, schoolRepo)
	quotationHandler := handler.NewQuotationHandler(quotationService)

	priceRefService := service.NewPriceRefService(priceRefRepo, userRepo)
	priceRefHandler := handler.NewPriceRefHandler(priceRefService)

	productionService := service.NewProductionService(productionRepo, userRepo, schoolRepo)
	productionHandler := handler.NewProductionHandler(productionService)

	api := r.Group("/api")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api.POST("/login", loginHandler.Login)

	api.Use(authMiddleware())
	userRoutes := api.Group("/user")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
		userRoutes.GET("/info", userHandler.GetInfoUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:userId", userHandler.UpdateUser)
		userRoutes.DELETE("/:userId", userHandler.DeleteUser)
	}
	schoolRoutes := api.Group("/school")
	{
		schoolRoutes.GET("/", schoolHandler.GetSchoolByUserId)
		schoolRoutes.POST("/", schoolHandler.CreateSchool)
	}
	quotationRoutes := api.Group("/quotation")
	{
		quotationRoutes.GET("/", quotationHandler.GetAllQuotation)
		quotationRoutes.POST("/", quotationHandler.CreateQuotation)
	}
	priceRefRoutes := api.Group("/priceRef")
	{
		priceRefRoutes.GET("/", priceRefHandler.GetPriceRefByUserID)
		priceRefRoutes.POST("/", priceRefHandler.CreatePriceRef)
	}
	productionRoutes := api.Group("/production")
	{
		productionRoutes.GET("/:productionId", productionHandler.GetProductionByID)
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			helper.ErrorResponse(c, http.StatusUnauthorized, helper.ErrMissingToken)
			return
		}
		splitToken := strings.Split(authorization, "Bearer ")
		tokenString := splitToken[1]
		if len(splitToken) < 2 {
			helper.ErrorResponse(c, http.StatusUnauthorized, helper.ErrMissingToken)
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			helper.ErrorResponse(c, http.StatusUnauthorized, helper.ErrInvalidToken)
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			helper.ErrorResponse(c, http.StatusUnauthorized, helper.ErrInvalidToken)
			return
		}

		c.Next()
	}
}
