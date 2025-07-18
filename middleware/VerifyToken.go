package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// VerifyToken is a middleware function that verifies the JWT token from the Authorization header.
// It checks the token's validity, expiration, and extracts the user ID from the token claims.
// If the token is valid, the user ID is stored in the request context for subsequent handlers to use.
// If the token is invalid or expired, it returns a 401 Unauthorized status with an appropriate error message.

func VerifyToken(ctx *fiber.Ctx) error {
	authToken := ctx.Get("Authorization")

	if authToken == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "did not receive a valid token",
		})
	}

	parsedToken, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("error ==>", err)
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
		userID := claims["userId"]
		ctx.Locals("userId", userID)

		return ctx.Next()
	}

	// Token is invalid
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "invalid token",
	})
}
