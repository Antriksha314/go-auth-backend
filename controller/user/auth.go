package user

import (
	"auth-backend/dto"
	"auth-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(response *gin.Context) {
	var payload dto.RegisterDTO

	if err := response.ShouldBind(&payload); err != nil {
		response.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	// var user := models.User.Find()

	// if user {
	// 	response.JSON(http.StatusConflict, gin.H{
	// 		"success": true,
	// 		"message": "User successfully register!",
	// 	})
	// 	return
	// }

	fmt.Printf("response data from regiter api -----> ", &payload)

	result := models.DB.Create(payload)

	if result.Error != nil {
		response.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"message": result.Error,
		})
	}

	response.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User successfully register!",
	})

}

func Login(response *gin.Context) {
	var payload dto.LoginDTO
	if err := response.BindJSON(&payload); err != nil {
		response.AbortWithStatus(400)
		return
	}
	fmt.Printf("response data from Login api -----> ", &payload)

}
