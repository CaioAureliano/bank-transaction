package authentication

import (
	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(configuration.Env.JWTSECRET),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}
