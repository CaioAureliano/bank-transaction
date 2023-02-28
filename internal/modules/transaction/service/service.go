package service

import (
	"errors"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
)

type repository interface {
	PubMessage(*domain.PubMessage) error
	CreateTransaction(t *domain.Transaction) (uint, error)
	ExistsByUserIDAndStatus(userID uint, status []domain.Status) bool
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

	message := &domain.PubMessage{
		Payer:         userID,
		Payee:         req.Payee,
		Value:         req.Value,
		TransactionID: transactionID,
	}

	if err := s.r.PubMessage(message); err != nil {
		return 0, err
	}

	return transactionID, nil
}
