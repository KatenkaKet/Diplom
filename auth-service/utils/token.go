package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, jwtKey string) (string, error) {
	log.Println(jwtKey)
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
