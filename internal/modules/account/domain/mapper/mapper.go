package mapper

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

func ToModel(d dto.CreateRequestDTO) *domain.User {
	return &domain.User{
		Firstname: d.Firstname,
		Lastname:  d.Lastname,
		Email:     d.Email,
		CPF:       d.CPF,
		Password:  d.Password,
		Account: &domain.Account{
			Type: domain.Type(d.Type),
		},
	}
}

func ToEntity(m *domain.User) *model.User {
	return &model.User{
		Firstname: m.Firstname,
		Lastname:  m.Lastname,
		Email:     m.Email,
		CPF:       m.CPF,
		Password:  m.Password,
		Account: &model.Account{
			Type: model.Type(m.Account.Type),
		},
	}
}
