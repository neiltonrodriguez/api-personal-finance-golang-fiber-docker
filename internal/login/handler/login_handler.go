package handler

import (
	"personal-finance-api/domain"
	AuthModel "personal-finance-api/models/auth"

	"github.com/gofiber/fiber/v2"
)

func Login(fiberCtx *fiber.Ctx) error {
	payload := domain.Login{}
	ctx := fiberCtx.Context()

	if err := fiberCtx.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	token, user, err := AuthModel.SignIn(ctx, payload.Email, payload.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Authorization{
		Token: token,
		Data:       user,
	})
}
