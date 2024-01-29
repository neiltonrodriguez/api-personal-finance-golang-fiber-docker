package handler

import (
	"personal-finance-api/domain"
	TransactionTypeModel "personal-finance-api/models/transaction_type"

	"github.com/gofiber/fiber/v2"
)

func Get(fiberCtx *fiber.Ctx) error {

	ctx := fiberCtx.Context()

	result, err := TransactionTypeModel.Get(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
		Meta: domain.Meta{
			Count: len(result),
		},
		Data: result,
	})
}

// func GetById(fiberCtx *fiber.Ctx) error {
// 	ctx := fiberCtx.Context()
// 	id, _ := strconv.Atoi(fiberCtx.Params("id"))

// 	result, err := TransactionModel.GetById(ctx, id)
// 	if err != nil {
// 		if err.Error() == "transaction not found" {
// 			return fiber.NewError(fiber.StatusNotFound, err.Error())
// 		}
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
// 		Meta: domain.Meta{
// 			Count: 1,
// 		},
// 		Data: result,
// 	})
// }

// func Create(fiberCtx *fiber.Ctx) error {
// 	payload := domain.TransactionInput{}
// 	ctx := fiberCtx.Context()

// 	if err := fiberCtx.BodyParser(&payload); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	result, err := TransactionModel.Create(ctx, payload)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
// 		Meta: domain.Meta{
// 			Count: 1,
// 		},
// 		Data: result,
// 	})
// }

// func Update(fiberCtx *fiber.Ctx) error {
// 	payload := domain.User{}
// 	ctx := fiberCtx.Context()

// 	if err := fiberCtx.BodyParser(&payload); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	id, _ := strconv.Atoi(fiberCtx.Params("id"))
// 	err := TransactionModel.Update(ctx, id, payload)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
// 		Meta: domain.Meta{
// 			Count: 1,
// 		},
// 		Data: "updated with success",
// 	})
// }

// func Delete(fiberCtx *fiber.Ctx) error {
// 	ctx := fiberCtx.Context()
// 	claims := fiberCtx.Locals("user").(*jwt.Token)
// 	mapClaims := claims.Claims.(jwt.MapClaims)
// 	data, err := json.Marshal(mapClaims["user"])
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	var user domain.User
// 	err = json.Unmarshal(data, &user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	id, _ := strconv.Atoi(fiberCtx.Params("id"))

// 	err = TransactionModel.Delete(ctx, id)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, err.Error())
// 	}

// 	return fiberCtx.Status(fiber.StatusOK).JSON(domain.Response{
// 		Meta: domain.Meta{
// 			Count: 1,
// 		},
// 		Data: "Deleted with success",
// 	})
// }
