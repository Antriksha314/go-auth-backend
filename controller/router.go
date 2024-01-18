package controller

import (
	"auth-backend/controller/user"

	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	router.POST("/register", user.Register)
	router.POST("/login", user.Login)
}
