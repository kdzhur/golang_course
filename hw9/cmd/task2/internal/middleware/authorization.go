package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Handler = func(*fiber.Ctx) error

const API_KEY = "teacher"

func Authorization() Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || authHeader != API_KEY {
			return fiber.ErrForbidden
		}
		return c.Next()
	}
}
