package handlers

import "github.com/gofiber/fiber/v2"

func GhLogin(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"msg": "Success",
	})
}
