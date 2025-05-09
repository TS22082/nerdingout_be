package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// VerifyToken verifies a token and returns a JSON response
func VerifyToken(ctx *fiber.Ctx) error {
	// @param ctx The Fiber context object.
	// @return nil if the verification was successful, or an error message otherwise.
	return ctx.JSON(fiber.Map{
		"success": true,
		"id":      ctx.Locals("userId"),
	})
}
