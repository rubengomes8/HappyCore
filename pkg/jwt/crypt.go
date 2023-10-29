package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", bcrypt.ErrMismatchedHashAndPassword
	}
	return string(hash), nil
}

func verifyPassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
