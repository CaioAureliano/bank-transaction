package repository

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/mapper"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) *gorm.DB
}

type Repository struct {
	db Database
}

func New(db Database) Repository {
	return Repository{db}
}

func (r Repository) Create(u *domain.User) error {
	return r.db.Create(mapper.ToEntity(u)).Error
}
