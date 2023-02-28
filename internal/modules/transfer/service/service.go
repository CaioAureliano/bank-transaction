package service

import (
	"log"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

type repository interface {
	GetAccountByID(uint) (*model.Account, error)
	CacheStatus(model.Status, uint) error
	Authenticator() error
	UpdateAccounts(*domain.Transference) error
	UpdateTransaction(*model.Transaction) error
}

type Service struct {
	r repository
}

func New(r repository) Service {
	return Service{r}
}

func (s Service) Transfer(message *domain.TransactionMessage) error {

	status, msg := s.process(message)

	log.Println(msg)

	defer s.r.CacheStatus(status, message.TransactionID)

	transaction := &model.Transaction{
		ID:      message.TransactionID,
		PayerID: message.Payer,
		PayeeID: message.Payee,
		Value:   message.Value,
		Status:  status,
		Message: msg,
	}

	return s.r.UpdateTransaction(transaction)
}

func (s Service) process(msg *domain.TransactionMessage) (model.Status, string) {

	if err := s.r.CacheStatus(model.PROCESSING, msg.TransactionID); err != nil {
		return model.FAILED, err.Error()
	}

	payer, _ := s.r.GetAccountByID(msg.Payer)
	payee, err := s.r.GetAccountByID(msg.Payee)
	if err != nil {
		return model.FAILED, err.Error()
	}

	if err := s.r.Authenticator(); err != nil {
		return model.FAILED, err.Error()
	}

	t := domain.NewTransference(payer, payee, msg.Value)
	if err := t.Transfer(); err != nil {
		return model.FAILED, err.Error()
	}

	if err := s.r.UpdateAccounts(t); err != nil {
		return model.FAILED, err.Error()
	}

	return model.SUCCESS, "successful transaction"
}
