package service

import (
	"errors"
	"strconv"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/CaioAureliano/bank-transaction/pkg/utils"
)

type repository interface {
	SendMessage(*domain.TransactionQueueMessage) error
	CreateTransaction(t *domain.Transaction) (uint, error)
	ExistsByUserIDAndStatus(userID uint, status []domain.Status) bool
	GetTransactionByIDAndPayerID(transactionID, payerID uint) (*model.Transaction, error)
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

func (s Service) GetTransaction(req *dto.GetTransactionRequestDTO) (*dto.TransactionResponseDTO, error) {
	transactionID, _ := strconv.ParseUint(req.TransactionID, 10, 64)
	transactionPersisted, err := s.r.GetTransactionByIDAndPayerID(uint(transactionID), req.PayerID)
	if err != nil {
		return nil, err
	}

	transaction := utils.ParseTo[domain.Transaction](transactionPersisted)

	return &dto.TransactionResponseDTO{
		Status:  transaction.Status,
		Message: transaction.Status.String(),
	}, nil
}
