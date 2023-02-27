package repository

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/mapper"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) *gorm.DB
}

type Broker interface{}

type Repository struct {
	db Database
	b  Broker
}

func New(db Database, b Broker) Repository {
	return Repository{db, b}
}

func (r Repository) CreateTransaction(t *domain.Transaction) (uint, error) {
	entity := mapper.ToEntity(t)
	result := r.db.Create(&entity)
	if result.Error != nil {
		return 0, result.Error
	}

	return entity.ID, nil
}

func (r Repository) ExistsByUserIDAndStatus(userID uint, status []domain.Status) bool {
	return false
}

func (r Repository) PubMessage(*domain.PubMessage) error {
	return nil
}
