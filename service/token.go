package service

import (
	err "github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type AuthService interface {
	GetToken() (*middleware.AuthToken, *err.Error)
}

type JwtAuthService struct {
}

var SECRET_KEY = os.Getenv("SECRET_KEY")

func (jas JwtAuthService) GetToken() (*middleware.AuthToken, *err.Error) {
	claims := middleware.CustomClaims{
		Authenticated: true,
		StandardClaims: jwt.StandardClaims{
			Issuer: os.Getenv("APP_NAME"),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, e := token.SignedString([]byte(SECRET_KEY))
	if e != nil {
		return nil, err.ServerError("token could not be created")
	}
	return &middleware.AuthToken{Token: signedToken}, nil
}
