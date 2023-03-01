package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/CaioAureliano/bank-transaction/pkg/authentication"
	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

type mockService struct {
	fnCreateTransaction func(*dto.TransactionRequestDTO, uint) (uint, error)
	fnGetTransaction    func(*dto.GetTransactionRequestDTO) *dto.TransactionResponseDTO
}

func (m mockService) CreateTransaction(req *dto.TransactionRequestDTO, userID uint) (uint, error) {
	if m.fnCreateTransaction == nil {
		return 0, nil
	}
	return m.fnCreateTransaction(req, userID)
}

func (m mockService) GetTransaction(req *dto.GetTransactionRequestDTO) *dto.TransactionResponseDTO {
	if m.fnGetTransaction == nil {
		return nil
	}
	return m.fnGetTransaction(req)
}

var getRequest = func(method, endpoint, token, body string) *http.Request {
	req := httptest.NewRequest(method, endpoint, bytes.NewBuffer([]byte(body)))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token)
	return req
}

var jwtMiddleware = func() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(configuration.Env.JWTSECRET),
		SigningMethod: jwt.SigningMethodHS256.Name,
	})
}

func TestCreateTransaction(t *testing.T) {
	t.Parallel()

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
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO, userID uint) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 1, time.Now().Add(time.Minute*1))
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
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO, userID uint) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 2, time.Now().Add(time.Minute*10))
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
				fnCreateTransaction: func(trd *dto.TransactionRequestDTO, userID uint) (uint, error) {
					return 0, nil
				},
			},
			jwtMock: func() string {
				t, _ := authentication.GenerateJwt(1, 2, time.Now().Add(time.Minute-10))
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
			app.Use(jwtMiddleware())

			Router(app, h)

			req := getRequest(fiber.MethodPost, transactionEndpoint, tt.jwtMock, tt.body)

			res, err := app.Test(req, -1)

			tt.expectError(t, err)
			assert.Equal(t, tt.expectStatusCode, res.StatusCode)
		})
	}
}

func TestGetTransaction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string

		transactionIDMock int
		serviceMock       service

		expectStatus   int
		expectResponse string
	}{
		{
			name: "should be return Ok status with valid request and service return",

			transactionIDMock: 2,
			serviceMock: mockService{
				fnGetTransaction: func(req *dto.GetTransactionRequestDTO) *dto.TransactionResponseDTO {
					return &dto.TransactionResponseDTO{
						Status:  domain.SUCCESS,
						Message: domain.SUCCESS.String(),
					}
				},
			},

			expectStatus:   fiber.StatusOK,
			expectResponse: `{"status":3,"message":"SUCCESS"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := api.Setup()
			app.Use(jwtMiddleware())

			h := New(tt.serviceMock)
			Router(app, h)

			jwtMock, _ := authentication.GenerateJwt(1, 1, time.Now().Add(time.Minute*10))
			req := getRequest(fiber.MethodGet, fmt.Sprintf("%s/%d", transactionEndpoint, tt.transactionIDMock), jwtMock, "")

			res, _ := app.Test(req, -1)
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.expectResponse, string(body))
		})
	}
}
