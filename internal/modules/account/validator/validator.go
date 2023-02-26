package validator

import (
	"errors"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
)

type repository interface {
	ExistsByCpfOrEmail(cpf, email string) bool
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
	return v.r.ExistsByCpfOrEmail(user.CPF, user.Email)
}

func (v Validator) isValidCPF(cpf string) bool {
	return true
}
