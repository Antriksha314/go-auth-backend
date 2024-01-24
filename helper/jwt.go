package helper

import (
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTtoken(email string) (*string, error, int) {
	var secretKey = []byte(os.Getenv("JWT_KEY"))

	data := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	token, err := data.SignedString(secretKey)

	if err != nil {
		return nil, errors.New("Something went wrong"), http.StatusInternalServerError
	}
	return &token, nil, http.StatusOK
}
func VerifyJWTtoken(token string) (*jwt.Token, error, int) {
	var secretKey = []byte(os.Getenv("JWT_KEY"))

	data, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("Invalid token"), http.StatusUnauthorized
	}

	return data, nil, http.StatusOK

}
