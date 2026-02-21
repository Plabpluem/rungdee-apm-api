package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/extractors"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c fiber.Ctx) error {
	extractor := extractors.FromAuthHeader("Bearer")
	accessToken, err := extractor.Extract(c)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.ParseWithClaims(accessToken, jwt.MapClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	claim := token.Claims.(jwt.MapClaims)
	c.Locals("payload", claim)
	c.Locals("role", claim["role"])
	return c.Next()
}
