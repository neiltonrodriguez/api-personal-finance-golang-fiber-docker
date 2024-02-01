package main

import (
	"personal-finance-api/domain"
	AccountRouter "personal-finance-api/internal/account/router"
	CategoryRouter "personal-finance-api/internal/category/router"
	CheckRouter "personal-finance-api/internal/check/router"
	LoginRouter "personal-finance-api/internal/login/router"
	PaymentTypeRouter "personal-finance-api/internal/payment_type/router"
	TransactionRouter "personal-finance-api/internal/transaction/router"
	TransactionTypeRouter "personal-finance-api/internal/transaction_type/router"
	UserRouter "personal-finance-api/internal/user/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	domain.GlobalConfig.LoadVariables()

	UserRouter.RegisterRoutes(app)
	AccountRouter.RegisterRoutes(app)
	CategoryRouter.RegisterRoutes(app)
	TransactionRouter.RegisterRoutes(app)
	TransactionTypeRouter.RegisterRoutes(app)
	PaymentTypeRouter.RegisterRoutes(app)
	LoginRouter.RegisterRoutes(app)
	CheckRouter.RegisterRoutes(app)
	app.Listen(":8080")
}
