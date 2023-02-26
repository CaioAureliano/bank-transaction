package service

import "github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"

type repository interface {
	Create(*domain.User) error
}

type validator interface {
	Validate(*domain.User) error
}
