package router

import (
	"personal-finance-api/domain"
	"personal-finance-api/internal/check/handler"
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
		"v1/check",
		middlewares.Restricted,
	)

	route.Use(cors.New())

	route.Use(config)

	route.Post("/", func(fiberCtx *fiber.Ctx) error {
		return handler.CheckToken(fiberCtx)
	})

}
