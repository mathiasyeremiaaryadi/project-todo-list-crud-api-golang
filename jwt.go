package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user User) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"name":   user.Name,
		"email":  user.Email,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenResult, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenResult, nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	verifiedToken, err := jwt.Parse(token, func(verifiedToken *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return verifiedToken, err
	}

	if !verifiedToken.Valid {
		return verifiedToken, fmt.Errorf("invalid token")
	}

	return verifiedToken, nil
}
