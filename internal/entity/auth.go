package entity

import "github.com/golang-jwt/jwt/v5"

type AuthToken struct {
	Token string `json:"token"`
}

type AuthClaim struct {
	jwt.RegisteredClaims
	Name  string `json:"name"`
	Email string `json:"email"`
}
