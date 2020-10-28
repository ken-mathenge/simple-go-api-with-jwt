package main

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

// CreateJWT function generates a signed JWT
func CreateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(), // Could have more payload
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("can't sign the token: %v", err)
	}

	return tokenString, nil
}
