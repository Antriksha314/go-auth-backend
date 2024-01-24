package controller

import (
	"auth-backend/dto"
	"auth-backend/models"
	"auth-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(response *gin.Context) {
	user := response.MustGet("user").(models.User)

	if user.Email == nil {
		response.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"success": false,
		})
	}
	response.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})

}

func Users(response *gin.Context) {
	users, err, code := services.AllUsersService(response)

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

func Update(response *gin.Context) {
	var payload dto.UpdateUser
	if err := response.BindJSON(&payload); err != nil {
		response.AbortWithStatus(400)
		return
	}
	_, err, code := services.UpdateUserService(response, payload)

	if err != nil {
		response.JSON(code, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response.JSON(code, gin.H{
		"success": true,
		"message": "User updated successfully",
	})
	return
}

func ChangePassword(response *gin.Context) {
	var payload dto.ChangePassword

	if err := response.BindJSON(&payload); err != nil {
		response.AbortWithStatus(400)
		return
	}

	_, err, code := services.ChangePassword(response, payload)

	if err != nil {
		response.JSON(code, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response.JSON(code, gin.H{
		"success": true,
		"message": "Password change successfully",
	})
	return
}
