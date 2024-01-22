package main

import (
	"personal-finance-api/domain"
	AccountRouter "personal-finance-api/internal/account/router"
	LoginRouter "personal-finance-api/internal/login/router"
	TransactionRouter "personal-finance-api/internal/transaction/router"
	UserRouter "personal-finance-api/internal/user/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	domain.GlobalConfig.LoadVariables()
	// app.Get("/", HelloHandler)

	UserRouter.RegisterRoutes(app)
	AccountRouter.RegisterRoutes(app)
	TransactionRouter.RegisterRoutes(app)
	LoginRouter.RegisterRoutes(app)
	app.Listen(":8080")
}

// func HelloHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World!sssss")
// }
