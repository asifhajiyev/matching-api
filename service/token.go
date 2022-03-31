package service

import (
	err "github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

type AuthService interface {
	GetToken() (*middleware.AuthToken, *err.Error)
}

type JwtAuthService struct {
}

var SecretKey = os.Getenv("SECRET_KEY")

func (jas JwtAuthService) GetToken() (*middleware.AuthToken, *err.Error) {
	claims := middleware.CustomClaims{
		Authenticated: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, e := token.SignedString([]byte(SecretKey))
	if e != nil {
		return nil, err.ServerError("token could not be created")
	}
	return &middleware.AuthToken{Token: signedToken}, nil
}
