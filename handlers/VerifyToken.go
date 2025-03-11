package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"os"
	"time"
)

func VerifyToken(ctx *fiber.Ctx) error {
	authToken := ctx.Get("Authorization")

	if authToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "did not receive a valid token",
		})
	}

	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "failed to parse token",
		})
	}

	// Validate token and extract claims
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		// Check the expiration time
		if exp, ok := claims["exp"].(float64); ok {
			if int64(exp) < time.Now().Unix() {
				return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "token is expired",
				})
			}
		} else {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token claims",
			})
		}

		// Token is valid and not expired
		userID := claims["id"]
		ctx.Locals("userId", userID)

		return ctx.Next()
	}

	// Token is invalid
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "invalid token",
	})
}
