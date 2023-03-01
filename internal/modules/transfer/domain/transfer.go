package domain

import (
	"errors"
	"fmt"

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

func (t *Transference) Transfer() error {
	if t.Payer.Balance < t.Value {
		return errors.New(fmt.Sprintf("payer(user account with ID: %d) insufficient balance", t.Payer.ID))
	}

	t.Payer.Balance -= t.Value
	t.Payee.Balance += t.Value

	return nil
}
