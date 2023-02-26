package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/authentication"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	fnCreateTransaction func(*dto.TransactionRequestDTO) (uint, error)
}

func (m mockService) CreateTransaction(req *dto.TransactionRequestDTO) (uint, error) {
	if m.fnCreateTransaction == nil {
		return 0, nil
	}
	return m.fnCreateTransaction(req)
}

func TestCreateTransaction(t *testing.T) {

	tests := []struct {
		name string

		serviceMock mockService
		jwtMock     string
		body        string

		expectError      assert.ErrorAssertionFunc
		expectStatusCode int
	}{
		{
			name: "should be return accepted(202) status with valid JWT",

			serviceMock: mockService{
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 0, time.Now().Add(time.Minute*1))
				return t
			}(),
			body: `{
				"value": 99.9,
				"payee": 2
			}`,

			expectError:      assert.NoError,
			expectStatusCode: fiber.StatusAccepted,
		},
		{
			name: "should be return unauthorized(401) status with invalid user type",

			serviceMock: mockService{
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 1, time.Now().Add(time.Minute*10))
				return t
			}(),
			body: `{
				"value": 99.9,
				"payee": 2
			}`,

			expectError:      assert.NoError,
			expectStatusCode: fiber.StatusUnauthorized,
		},
		{
			name: "should be return unauthorized(401) status with invalid jwt expires",

			serviceMock: mockService{
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 1, time.Now().Add(time.Minute-10))
				return t
			}(),
			body: `{
				"value": 99.9,
				"payee": 2
			}`,

			expectError:      assert.NoError,
			expectStatusCode: fiber.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			h := New(tt.serviceMock)
			app := api.Setup()
			app.Use(jwtware.New(jwtware.Config{
				SigningKey:    []byte(authentication.JWT_SECRET),
				SigningMethod: jwt.SigningMethodHS256.Name,
			}))

			Router(app, h)

			req := httptest.NewRequest(fiber.MethodPost, transactionEndpoint, bytes.NewBuffer([]byte(tt.body)))
			req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			req.Header.Set(fiber.HeaderAuthorization, "Bearer "+tt.jwtMock)

			res, err := app.Test(req, -1)

			tt.expectError(t, err)
			assert.Equal(t, tt.expectStatusCode, res.StatusCode)
		})
	}
}
