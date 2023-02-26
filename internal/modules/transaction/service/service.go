package service

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
)

type repository interface {
	PubMessage(*domain.PubMessage) error
	CreateTransaction(t *domain.Transaction) (uint, error)
}

type Service struct {
	r repository
}

func New(r repository) Service {
	return Service{r}
}

func (s Service) CreateTransaction(req *dto.TransactionRequestDTO, userID uint) (uint, error) {

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
		Payer: userID,
		Payee: req.Payee,
		Value: req.Value,
	}

	if err := s.r.PubMessage(message); err != nil {
		return 0, err
	}

	return transactionID, nil
}
