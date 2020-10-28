package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var secret = []byte(os.Getenv("JWT_SECRET"))

// AuthenticationMiddleware handles jwt authentication
func AuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header["Authorization"] != nil {
				tokenString := r.Header["Authorization"][0]
				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return secret, nil
				})
				if err != nil {
					fmt.Fprintf(w, err.Error())
				}
				if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					next.ServeHTTP(w, r)
					return
				}
			}
			fmt.Fprintf(w, "Unauthorized: %v", http.StatusUnauthorized)
		})
	}
}
