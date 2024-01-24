package middleware

import (
	"auth-backend/helper"
	"auth-backend/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Protect(response *gin.Context) {
	authorization := response.Request.Header.Get("Authorization")
	if authorization == "" {
		response.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "You need to login to visit this route",
		})
		response.Abort()
		return
	}

	separateBearer := strings.Replace(authorization, "Bearer", "", 1)

	token := strings.TrimSpace(separateBearer)

	data, err, code := helper.VerifyJWTtoken(token)

	if err != nil {
		response.JSON(code, gin.H{
			"success": false,
			"message": err.Error(),
		})
		response.Abort()
		return
	}
	result := data.Claims.(jwt.MapClaims)

	if !data.Valid {
		response.JSON(code, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		response.Abort()
		return
	}
	email := result["email"]

	// Convert email type of interface to string
	emailString := fmt.Sprintf("%v", email)
	user, err, code := services.FindUserService(response, emailString)

	if err != nil {
		response.JSON(code, gin.H{
			"success": false,
			"message": err.Error(),
		})
		response.Abort()
		return
	}
	response.Set("user", user)
	response.Next()
}
