package controller

import (
	"auth-backend/dto"
	"auth-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(response *gin.Context) {
	var payload dto.RegisterDTO

	if error := response.ShouldBind(&payload); error != nil {
		response.JSON(http.StatusBadRequest, gin.H{
			"message": error.Error(),
			"success": false,
		})
		return
	}

	_, err, code := services.RegisterUserService(response, payload)

	if err != nil {
		response.JSON(code, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	response.JSON(code, gin.H{
		"message": "User successfully register!",
		"success": true,
	})

}

func Login(response *gin.Context) {
	var payload dto.LoginDTO
	if err := response.BindJSON(&payload); err != nil {
		response.AbortWithStatus(400)
		return
	}

	token, err, code := services.LoginUserService(response, payload)
	if err != nil {
		response.JSON(code, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	response.JSON(code, gin.H{
		"message": "Login successfully!",
		"success": true,
		"data":    token,
	})

}
