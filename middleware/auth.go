package middleware

import (
	"errors"
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

type CustomClaims struct {
	Authenticated bool `json:"authenticated"`
	jwt.RegisteredClaims
}

type AuthToken struct {
	Token string `json:"token"`
}

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:     []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler:   jwtError,
		SuccessHandler: jwtSuccess,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(response.RestResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(response.RestResponse{
		Code:    http.StatusUnauthorized,
		Message: err.Error(),
		Data:    nil,
	})
}

func jwtSuccess(c *fiber.Ctx) error {
	err := verifyToken(c)
	if err != nil {
		return jwtError(c, err)
	}
	return c.Next()
}

func extractToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is required")
	}

	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
		return "", errors.New("wrong authorization header")
	}
	return bearerToken[1], nil
}

func verifyToken(c *fiber.Ctx) error {
	tokenString, err := extractToken(c)

	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(
		tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		},
	)
	if err != nil {
		return err
	}

	claims := token.Claims.(*CustomClaims)
	if !claims.Authenticated {
		return errors.New("invalid token")
	}
	return nil
}
