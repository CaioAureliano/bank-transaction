package service

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

type repository interface {
	Create(*domain.User) error
	GetByEmail(string) (*model.User, error)
}

type validator interface {
	Validate(*domain.User) error
}
