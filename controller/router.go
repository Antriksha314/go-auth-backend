package controller

import (
	"auth-backend/middleware"

	"github.com/gin-gonic/gin"
)

func Routers(router *gin.RouterGroup) {
	router.POST("/register", Register)
	router.POST("/login", Login)

	// Protected routes
	router.Use(middleware.Protect)
	router.GET("/user", User)
	router.GET("/users", Users)
	router.PUT("/change-password", ChangePassword)
	router.PUT("/user/update", Update)
}
