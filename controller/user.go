package controller

import (
	"auth-backend/services"

	"github.com/gin-gonic/gin"
)

func User(response *gin.Context) {
	email := response.Param("email")
	user, err, code := services.FindUser(response, email)

	if err != nil {
		response.JSON(code, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	response.JSON(code, gin.H{
		"success": true,
		"data":    user,
	})

}

func Users(response *gin.Context) {
	users, err, code := services.AllUsers(response)

	if err != nil {
		response.JSON(code, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	response.JSON(code, gin.H{
		"success": true,
		"data":    users,
	})

}
