package main

import (
	"log"
	"notebook-backend/config"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//	@title		Tickbook API
//	@version	1.0
// 	@termsOfService http://tickbook.net/

// @contact.name API Support
// @contact.url http://tickbook.net/support
// @contact.email support@tickbook.net

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@schemes	https http

// @securityDefinitions.apikey	JwtToken
// @in							header
// @name						Authorization
// @BasePath 				/api

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// setup env
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	// database setup
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	db := config.InitDB(dbUsername, dbPassword, dbName, dbHost, dbPort)

	// setup routes
	r := gin.Default()

	r.Use(cors.Default())
	config.SetupRoutes(r, db)

	r.Run()
}
