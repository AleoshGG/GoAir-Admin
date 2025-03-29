package entities

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Password string    `json:"password"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}