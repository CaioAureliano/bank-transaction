package domain

import (
	"errors"

	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

type Transference struct {
	Payer *model.Account
	Payee *model.Account
	Value float64
}

func NewTransference(payer, payee *model.Account, value float64) *Transference {
	return &Transference{
		Payer: payer,
		Payee: payee,
		Value: value,
	}
}

var (
	ErrInvalidUserType = errors.New("invalid user account type")
	ErrWithouBalance   = errors.New("payer insufficient balance")
)

func (t *Transference) Transfer() error {
	if t.Payer.Type != model.USER {
		return ErrInvalidUserType
	}

	if t.Payer.Balance == 0 || t.Payer.Balance < t.Value {
		return ErrWithouBalance
	}

	t.Payer.Balance -= t.Value
	t.Payee.Balance += t.Value

	return nil
}
