package main

import (
	"log"
	"notebook-backend/config"
	"notebook-backend/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	db := config.InitDB(dbUsername, dbPassword, dbName, dbHost, dbPort)

	r := gin.Default()

	// Setup routes
	router.SetupRoutes(r, db)

	r.Run()
}
