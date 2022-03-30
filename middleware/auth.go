package middleware

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Authenticated bool `json:"authenticated"`
	jwt.StandardClaims
}

type AuthToken struct {
	Token string `json:"token"`
}
