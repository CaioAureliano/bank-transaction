package service

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/mapper"
)

type Service struct {
	r repository
	v validator
}

func New(r repository, v validator) Service {
	return Service{r, v}
}

func (s Service) CreateUserAccount(req dto.CreateRequestDTO) error {

	account := mapper.ToModel(req)

	if err := s.v.Validate(&account.User); err != nil {
		return err
	}

	account.User.GeneratePassword()

	if err := s.r.Create(account); err != nil {
		return err
	}

	return nil
}
