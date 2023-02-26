package service

import (
	"log"

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

	user := mapper.ToModel(req)

	if err := s.v.Validate(user); err != nil {
		log.Printf("error to validate user - %s", err)
		return err
	}

	user.GeneratePassword()

	if err := s.r.Create(user); err != nil {
		log.Printf("error to try create user - %s", err)
		return err
	}

	return nil
}
