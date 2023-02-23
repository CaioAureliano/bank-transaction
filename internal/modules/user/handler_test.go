package user

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/shared/application"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	app := application.Setup()
	Router(app)

	body := CreateRequestDTO{
		Firstname: "User",
		Lastname:  "Example",
		Email:     "example@mail.com",
		CPF:       "0",
		Password:  "test1234",
		Type:      "user",
	}
	bodyJson, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/user", bytes.NewBuffer(bodyJson))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, res.StatusCode)
}
