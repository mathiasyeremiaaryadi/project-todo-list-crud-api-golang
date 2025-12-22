package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(userId int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"type":   "access",
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	})

	tokenResult, err := jwtToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenResult, nil
}

func GenerateRefreshToken(userId int) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"type":   "refresh",
		"exp":    time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	tokenResult, err := jwtToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenResult, nil
}

func VerifyToken(token string, secret string, expectedType string) (*jwt.Token, error) {
	verifiedToken, err := jwt.Parse(token, func(verifiedToken *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return verifiedToken, err
	}

	if !verifiedToken.Valid {
		return verifiedToken, fmt.Errorf("invalid token")
	}

	claims := verifiedToken.Claims.(jwt.MapClaims)
	tokenType := claims["type"].(string)
	if expectedType != tokenType {
		return verifiedToken, fmt.Errorf("invalid token type")
	}

	return verifiedToken, nil
}
