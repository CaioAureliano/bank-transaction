package authentication

import (
	"time"

	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwt(id, t uint, expiresAt time.Time) (string, error) {

	claims := jwt.MapClaims{
		"ID":   id,
		"type": t,
		"exp":  expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configuration.Env.JWTSECRET))
}
