package handler

import (
	"encoding/json"
	"personal-finance-api/domain"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken(c *fiber.Ctx) error {
	var user domain.User

	context := c.Locals("user").(*jwt.Token)
	claims := context.Claims.(jwt.MapClaims)
	dataUser, err := json.Marshal(claims["user"])
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataUser, &user)
	if err != nil {
		return err
	}

	isTrue := false
	if user.Id != 0 {
		isTrue = true
	}

	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(domain.ResponseCheck{
		Check: isTrue,
		User:  user,
	})
}
