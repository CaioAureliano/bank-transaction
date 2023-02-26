package handler

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/api"
	"github.com/gofiber/fiber/v2"
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

	serviceMock := mockService{
		func(trd *dto.TransactionRequestDTO) (uint, error) {
			return 0, nil
		},
	}

	h := New(serviceMock)
	app := api.Setup()

	Router(app, h)

	body := `{
		"value": 99.9,
		"payer": 1,
		"payee": 2
	}`

	req := httptest.NewRequest(fiber.MethodPost, transactionEndpoint, bytes.NewBuffer([]byte(body)))
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	res, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusAccepted, res.StatusCode)
}
