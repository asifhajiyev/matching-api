package services

import (
	"github.com/asifhajiyev/matching-api/error"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/asifhajiyev/matching-api/util"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type AuthService interface {
	GetToken() (*AuthToken, *error.Error)
}

type JwtAuthService struct {
}

type AuthToken struct {
	Token string `json:"token"`
}

func (jas JwtAuthService) GetToken() (*AuthToken, *error.Error) {
	secretKey := os.Getenv("SECRET_KEY")
	expireTime, err := util.StringToInt(os.Getenv("JWT_EXPIRE_TIME"))

	if err != nil {
		return nil, err
	}

	claims := middleware.CustomClaims{
		Authenticated: true,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("APP_NAME"),
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Minute * time.Duration(expireTime))},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, e := token.SignedString([]byte(secretKey))
	if e != nil {
		return nil, error.ServerError("token could not be created")
	}
	return &AuthToken{Token: signedToken}, nil
}
