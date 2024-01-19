package controller

import (
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.GET("/users", Users)
	router.GET("/user/:email", User)
}
