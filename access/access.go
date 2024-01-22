package access

import (
	"personal-finance-api/domain"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

var config = jwtware.New(jwtware.Config{
	SigningKey: jwtware.SigningKey{Key: []byte(domain.GlobalConfig.JwtSecretKey)}})

func Access() fiber.Handler {
	return config
}
