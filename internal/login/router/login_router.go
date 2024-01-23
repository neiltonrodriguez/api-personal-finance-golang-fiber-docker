package router

import (
	"personal-finance-api/internal/login/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(
	app *fiber.App,

) {

	route := app.Group(
		"v1/login",
	)

	route.Use(cors.New())

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Login(fiberCtx)
	})

}
