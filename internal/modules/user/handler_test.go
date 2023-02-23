package user

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/shared/application"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	tests := []struct {
		name string

		body string

		expectedStatusCode int
	}{
		{
			name: "should be return 201 Created status code with valid body",

			body: `{
				"firstname": "Abc",
				"lastname": "test",
				"email": "example@mail.com",
				"cpf": "0",
				"password": "test1234",
				"type": "user"
			}`,

			expectedStatusCode: fiber.StatusCreated,
		},
		{
			name: "should be return 422 Unprocessable Entity status code with invalid structure body",

			body: `{
				"key": "value",
			}`,

			expectedStatusCode: fiber.StatusUnprocessableEntity,
		},
		{
			name: "should be return 400 Unprocessable Entity status code with invalid value body",

			body: `{
				"firstname": "Abc",
				"lastname": "test",
				"email": "example@mail.com",
				"cpf": "0",
				"password": "test",
				"type": "user"
			}`,

			expectedStatusCode: fiber.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			app := application.Setup()
			Router(app)

			req := httptest.NewRequest(fiber.MethodPost, USER_ENDPOINT, bytes.NewBuffer([]byte(tt.body)))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			res, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)
		})
	}

}
