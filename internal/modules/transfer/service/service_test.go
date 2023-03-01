package service

import (
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	fnGetAccountByID    func(uint) (*model.Account, error)
	fnUpdateTransaction func(*model.Transaction) error
}

func (m mockRepository) GetAccountByID(id uint) (*model.Account, error) {
	if m.fnGetAccountByID == nil {
		return nil, nil
	}
	return m.fnGetAccountByID(id)
}

func (m mockRepository) CacheStatus(model.Status, uint) error {
	return nil
}

func (m mockRepository) Authenticator() error {
	return nil
}

func (m mockRepository) UpdateAccounts(t *domain.Transference) error {
	return nil
}

func (m mockRepository) UpdateTransaction(t *model.Transaction) error {
	if m.fnUpdateTransaction == nil {
		return nil
	}
	return m.fnUpdateTransaction(t)
}

func TestTransfer(t *testing.T) {

	tests := []struct {
		name string

		getAccountByIDMock func(u uint) (*model.Account, error)
		msgMock            *domain.TransactionMessage

		expectStatus model.Status
	}{
		{
			name: "should be SUCCESS status with valid values",

			getAccountByIDMock: func(u uint) (*model.Account, error) {
				return &model.Account{
					ID:      u + 2,
					UserID:  u,
					Balance: 100,
					Type:    model.USER,
				}, nil
			},
			msgMock: &domain.TransactionMessage{
				Value:         10,
				Payer:         1,
				Payee:         2,
				TransactionID: 100,
			},

			expectStatus: model.SUCCESS,
		},
		{
			name: "should be FAILED status with invalid balance and request value",

			getAccountByIDMock: func(u uint) (*model.Account, error) {
				return &model.Account{
					ID:      u + 2,
					UserID:  u,
					Balance: 5,
					Type:    model.USER,
				}, nil
			},
			msgMock: &domain.TransactionMessage{
				Value:         50,
				Payer:         1,
				Payee:         2,
				TransactionID: 100,
			},

			expectStatus: model.FAILED,
		},
		{
			name: "should be FAILED status with invalid account type",

			getAccountByIDMock: func(u uint) (*model.Account, error) {
				return &model.Account{
					ID:      u + 2,
					UserID:  u,
					Balance: 100,
					Type:    model.SELLER,
				}, nil
			},
			msgMock: &domain.TransactionMessage{
				Value:         10,
				Payer:         1,
				Payee:         2,
				TransactionID: 100,
			},

			expectStatus: model.FAILED,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			transactionSpy := new(model.Transaction)

			repositoryMock := mockRepository{
				fnGetAccountByID: tt.getAccountByIDMock,

				fnUpdateTransaction: func(transaction *model.Transaction) error {
					transactionSpy = transaction
					return nil
				},
			}

			s := New(repositoryMock)

			s.Transfer(tt.msgMock)

			assert.Equal(t, tt.expectStatus, transactionSpy.Status)
		})
	}

}
