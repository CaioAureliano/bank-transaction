package repository

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/mapper"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
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

func (r Repository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if result := r.db.First(&user, "email = ?", email); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
