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
	routers := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}
	routers.Use(cors.New(corsConfig))
	router := routers.Group("/api")
	controller.Routers(router)

	return routers
}
