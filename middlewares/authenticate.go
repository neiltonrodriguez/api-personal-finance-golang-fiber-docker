package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Accessible(c *fiber.Ctx) error {
	return c.Next()
}

func Restricted(c *fiber.Ctx) error {
	return c.Next()
}


