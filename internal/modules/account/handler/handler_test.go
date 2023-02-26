package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	fnCreateUserAccount func(dto.CreateRequestDTO) error
	fnAuthenticate      func(dto.AuthRequestDTO) (string, error)
}

func (m mockService) CreateUserAccount(req dto.CreateRequestDTO) error {
	if m.fnCreateUserAccount == nil {
		return nil
	}
	return m.fnCreateUserAccount(req)
}

func (m mockService) Authenticate(req dto.AuthRequestDTO) (string, error) {
	if m.fnAuthenticate == nil {
		return "", nil
	}
	return m.fnAuthenticate(req)
}

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
				"type": 1
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
				"type": 1
			}`,

			expectedStatusCode: fiber.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serviceMock := mockService{
				fnCreateUserAccount: func(crd dto.CreateRequestDTO) error {
					return nil
				},
			}

			h := New(serviceMock)
			app := api.Setup()

			Router(app, h)

			req := httptest.NewRequest(fiber.MethodPost, accountEndpoint, bytes.NewBuffer([]byte(tt.body)))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

			res, _ := app.Test(req, -1)

			assert.Equal(t, tt.expectedStatusCode, res.StatusCode)
		})
	}

}

func TestAuthenticate(t *testing.T) {

	serviceMock := mockService{
		fnAuthenticate: func(ard dto.AuthRequestDTO) (string, error) {
			return "validtoken", nil
		},
	}

	h := New(serviceMock)
	app := api.Setup()

	Router(app, h)

	body := `{
		"email": "example@mail.com",
		"password": "test1234"
	}`

	req := httptest.NewRequest(fiber.MethodPost, accountEndpoint+"/auth", bytes.NewBuffer([]byte(body)))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}
