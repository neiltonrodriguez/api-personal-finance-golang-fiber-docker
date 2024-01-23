package router

import (
	"personal-finance-api/domain"
	"personal-finance-api/internal/account/handler"
	"personal-finance-api/middlewares"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RegisterRoutes(
	app *fiber.App,

) {
	config := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(domain.GlobalConfig.JwtSecretKey)}})

	route := app.Group(
		"v1/account",
		middlewares.Restricted,
	)

	route.Use(cors.New())

	route.Use(config)

	route.Get("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Get(fiberCtx)
	})

	route.Get("/balance-total", func(fiberCtx *fiber.Ctx) error {
		return handler.GetBalanceTotal(fiberCtx)
	})

	route.Get("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.GetById(fiberCtx)
	})

	route.Put("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Update(fiberCtx)
	})

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.Create(fiberCtx)
	})

	route.Delete("/:id", func(fiberCtx *fiber.Ctx) error {
		return handler.Delete(fiberCtx)
	})

}
