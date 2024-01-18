package services

import (
	"auth-backend/dto"
	"auth-backend/helper"
	"auth-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserService(c *gin.Context, input dto.RegisterDTO) (*gorm.DB, error) {

	password, err := helper.ConvertToHashPassword(input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Something went wrong, please try again",
		})
	}
	user := models.User{FirstName: input.FirstName, LastName: input.LastName, Email: &input.Email, Password: password}

	result := models.DB.Create(&user)
	return result, nil
}

func LoginUserService(c *gin.Context, input dto.LoginDTO) (*gorm.DB, error) {

}
