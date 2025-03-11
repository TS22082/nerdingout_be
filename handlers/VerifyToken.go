package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func VerifyToken(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"success": true,
		"id":      ctx.Locals("userId"),
	})
}
