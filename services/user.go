package services

import (
	"auth-backend/dto"
	"auth-backend/helper"
	"auth-backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindUserService(c *gin.Context, email string) (models.User, error, int) {
	var user models.User
	result := models.DB.Find(&user, "email = ?", email)
	if result.Error != nil || user.Email == nil {
		return user, errors.New("User not found"), http.StatusNotFound
	}
	return user, nil, http.StatusOK
}

func AllUsersService(c *gin.Context) ([]models.User, error, int) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, errors.New("No record found"), http.StatusNotFound
	}
	return users, nil, http.StatusOK
}

func UpdateUserService(c *gin.Context, input dto.UpdateUser) (*gorm.DB, error, int) {
	user, err, code := FindUserService(c, input.Email)

	if err != nil {
		return nil, err, code
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	result := models.DB.Save(&user)

	return result, nil, http.StatusOK
}

func ChangePassword(c *gin.Context, input dto.ChangePassword) (*gorm.DB, error, int) {
	user, err, code := FindUserService(c, input.Email)

	if err != nil {
		return nil, err, code
	}

	passwordCheck := helper.ComparePassword(input.OldPassword, user.Password)

	if passwordCheck != nil {
		return nil, errors.New("Wrong Password"), http.StatusUnauthorized
	}

	hashPassword, err := helper.ConvertToHashPassword(input.Password)

	if err != nil {
		return nil, errors.New("Something went wrong,please try again"), http.StatusBadGateway
	}

	user.Password = hashPassword

	result := models.DB.Save(&user)

	return result, nil, http.StatusOK

}
