package model

import "github.com/dgrijalva/jwt-go"

// JWT claim payload
type JWTClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}
