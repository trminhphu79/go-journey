package utils

import (
	hasher "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := hasher.GenerateFromPassword([]byte(password), hasher.DefaultCost)
	return string(bytes), err
}

func ComparePassword(hash string, password string) bool {
	err := hasher.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
