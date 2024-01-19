package services

import (
	"auth-backend/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUser(c *gin.Context, email string) (models.User, error, int) {
	var user models.User
	result := models.DB.Find(&user, "email = ?", email)
	if result.Error != nil || user.Email == nil {
		return user, errors.New("User not found"), http.StatusNotFound
	}
	return user, nil, http.StatusOK
}

func AllUsers(c *gin.Context) ([]models.User, error, int) {
	var users []models.User
	result := models.DB.Find(&users)
	if result.Error != nil {
		return nil, errors.New("No record found"), http.StatusNotFound
	}
	return users, nil, http.StatusOK
}
