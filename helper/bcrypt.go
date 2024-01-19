package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func ConvertToHashPassword(value string) ([]byte, error) {
	password := []byte(value)
	result, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ComparePassword(password string, hashPassword []byte) error {
	error := bcrypt.CompareHashAndPassword(hashPassword, []byte(password))
	if error != nil {
		return error
	}
	return nil
}
