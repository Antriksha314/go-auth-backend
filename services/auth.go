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

func RegisterUserService(c *gin.Context, input dto.RegisterDTO) (*gorm.DB, error, int) {

	user, err, code := FindUser(c, input.Email)

	if err != nil && user.Email != nil {
		return nil, err, code
	}

	if user.Email != nil {
		return nil, errors.New("Email already exits"), http.StatusConflict
	}

	password, err := helper.ConvertToHashPassword(input.Password)

	if err != nil {
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}
	createUser := models.User{FirstName: input.FirstName, LastName: input.LastName, Email: &input.Email, Password: password}

	result := models.DB.Create(&createUser)
	return result, nil, http.StatusOK
}

func LoginUserService(c *gin.Context, input dto.LoginDTO) (*string, error, int) {
	user, err, code := FindUser(c, input.Email)

	if err != nil {
		return nil, err, code
	}

	passwordCheck := helper.ComparePassword(input.Password, user.Password)

	if passwordCheck != nil {
		return nil, errors.New("Wrong password"), http.StatusUnauthorized
	}
	return nil, nil, http.StatusOK
}
