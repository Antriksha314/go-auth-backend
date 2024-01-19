package main

import (
	"auth-backend/controller"
	"auth-backend/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	// Connection to database
	models.ConnectToDatabase()

	//Basic router
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}
	router.Use(cors.New(corsConfig))
	controller.Routers(router)

	return router
}
