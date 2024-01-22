package router

import (
	"personal-finance-api/internal/login/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(
	app *fiber.App,

) {

	route := app.Group(
		"v1/login",
	)

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Login(fiberCtx)
	})

}
