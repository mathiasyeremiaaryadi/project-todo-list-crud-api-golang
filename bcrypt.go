package main

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", nil
	}

	return string(passwordBytes), nil
}

func VerifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
