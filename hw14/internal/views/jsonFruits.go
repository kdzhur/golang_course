package views

import (
	"github.com/gofiber/fiber/v2"
)

func JsonFruits(c *fiber.Ctx, data ...any) error {
	return c.JSON(data)
}
