package entity

import "github.com/golang-jwt/jwt/v5"

type UserAuth struct {
	Token string `json:"token"`
}

type AuthClaim struct {
	jwt.RegisteredClaims
	Name  string `json:"name"`
	Email string `json:"Email"`
}
