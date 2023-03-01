package service

import (
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	fnSendMessage             func(*domain.TransactionQueueMessage) error
	fnCreateTransaction       func(*domain.Transaction) (uint, error)
	fnExistsByUserIDAndStatus func(userID uint, status []domain.Status) bool
}

func (m mockRepository) SendMessage(message *domain.TransactionQueueMessage) error {
	if m.fnSendMessage == nil {
		return nil
	}
	return m.fnSendMessage(message)
}

func (m mockRepository) CreateTransaction(t *domain.Transaction) (uint, error) {
	if m.fnCreateTransaction == nil {
		return 0, nil
	}
	return m.fnCreateTransaction(t)
}

func (m mockRepository) ExistsByUserIDAndStatus(userID uint, status []domain.Status) bool {
	if m.fnExistsByUserIDAndStatus == nil {
		return false
	}
	return m.fnExistsByUserIDAndStatus(userID, status)
}

func (m mockRepository) GetCachedStatusTransactionByID(transactionID string) *domain.Status {
	return nil
}

func TestCreateTransaction(t *testing.T) {

	repositoryMock := mockRepository{
		fnCreateTransaction: func(t *domain.Transaction) (uint, error) {
			return 111, nil
		},
	}

	userIdMock := uint(1)
	reqMock := &dto.TransactionRequestDTO{
		Value: 10000,
		Payee: 2,
	}

	s := New(repositoryMock)
	transactionID, err := s.CreateTransaction(reqMock, userIdMock)

	assert.NoError(t, err)
	assert.NotEmpty(t, transactionID)
}
