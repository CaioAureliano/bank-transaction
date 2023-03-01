package repository

import (
	"errors"
	"fmt"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/mapper"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
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
	user := new(model.User)
	if err := r.db.Joins("Account").First(&user, "email = ?", email).Error; err != nil || user == nil {
		return nil, errors.New(fmt.Sprintf("not found user with email \"%s\": %s", email, err.Error()))
	}
	return user, nil
}

func (r Repository) ExistsByCpfOrEmail(cpf, email string) bool {
	user := new(model.User)
	err := r.db.First(&user, "cpf = ? OR email = ?", cpf, email).Error
	return err == nil && user != nil
}
