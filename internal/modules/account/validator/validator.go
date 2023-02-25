package validator

import (
	"errors"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
)

type repository interface {
	Exists(*domain.User) bool
}

type Validator struct {
	r repository
}

func New(r repository) Validator {
	return Validator{r}
}

var (
	errUserExists = errors.New("user already exists")
	errInvalidCPF = errors.New("invalid cpf")
)

func (v Validator) Validate(u *domain.User) error {
	if v.ExistsByCpfOrEmail(u) {
		return errUserExists
	}

	if !v.isValidCPF(u.CPF) {
		return errInvalidCPF
	}

	return nil
}

func (v Validator) ExistsByCpfOrEmail(user *domain.User) bool {
	return v.r.Exists(user)
}

func (v Validator) isValidCPF(cpf string) bool {
	return true
}
