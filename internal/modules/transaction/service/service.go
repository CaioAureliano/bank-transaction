package service

import (
	"errors"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
)

type repository interface {
	SendMessage(*domain.TransactionQueueMessage) error
	CreateTransaction(t *domain.Transaction) (uint, error)
	ExistsByUserIDAndStatus(userID uint, status []domain.Status) bool
	GetCachedStatusTransactionByID(transactionID string) *domain.Status
}

type Service struct {
	r repository
}

func New(r repository) Service {
	return Service{r}
}

func (s Service) CreateTransaction(req *dto.TransactionRequestDTO, userID uint) (uint, error) {

	transactionCreatedStatus := []domain.Status{domain.REQUESTED, domain.PROCESSING}
	if s.r.ExistsByUserIDAndStatus(userID, transactionCreatedStatus) {
		return 0, errors.New("transaction in process: only one transaction by user")
	}

	transactionRequested := &domain.Transaction{
		PayerID: userID,
		PayeeID: req.Payee,
		Value:   req.Value,
		Status:  domain.REQUESTED,
	}

	transactionID, err := s.r.CreateTransaction(transactionRequested)
	if err != nil {
		return 0, err
	}

	message := &domain.TransactionQueueMessage{
		Payer:         userID,
		Payee:         req.Payee,
		Value:         req.Value,
		TransactionID: transactionID,
	}

	if err := s.r.SendMessage(message); err != nil {
		return 0, err
	}

	return transactionID, nil
}

func (s Service) GetTransaction(req *dto.GetTransactionRequestDTO) *dto.TransactionResponseDTO {

	status := s.r.GetCachedStatusTransactionByID(req.TransactionID)

	return &dto.TransactionResponseDTO{
		Status:  *status,
		Message: status.String(),
	}
}
