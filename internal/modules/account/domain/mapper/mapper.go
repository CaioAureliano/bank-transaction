package mapper

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
)

func ToModel(d dto.CreateRequestDTO) *domain.Account {
	return &domain.Account{
		User: &domain.User{
			Firstname: d.Firstname,
			Lastname:  d.Lastname,
			Email:     d.Email,
			CPF:       d.CPF,
			Password:  d.Password,
		},
		Type: domain.Type(d.Type),
	}
}
