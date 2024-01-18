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

func ComparePassword(password []byte, hashPassword []byte) (bool, error) {
	error := bcrypt.CompareHashAndPassword(hashPassword, password)
	if error != nil {
		return false, error
	}
	return true, nil
}
