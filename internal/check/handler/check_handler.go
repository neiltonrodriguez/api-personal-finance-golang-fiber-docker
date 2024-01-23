package handler

import (
	"personal-finance-api/domain"

	"github.com/gofiber/fiber/v2"
)

func CheckToken(fiberCtx *fiber.Ctx) error {
	// ctx := fiberCtx.Context()

	// _, err := AuthModel.CheckToken(ctx)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	// }

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.ResponseCheck{
		Check: true,
		Data:  "success",
	})
}
