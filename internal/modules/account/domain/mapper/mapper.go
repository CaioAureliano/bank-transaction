package mapper

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/CaioAureliano/bank-transaction/pkg/utils"
)

func RequestToModel(d dto.CreateRequestDTO) *domain.User {
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

func ToEntity(user *domain.User) *model.User {
	return utils.ParseTo[model.User](user)
}

func ToModel(entity *model.User) *domain.User {
	return utils.ParseTo[domain.User](entity)
}
