package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain/mapper"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Database interface {
	Create(interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}

type Queue interface {
	SendMessage(string) error
}

type Cache interface {
	Set(ctx context.Context, k string, v interface{}, expiration time.Duration) *redis.StatusCmd
}

type Repository struct {
	db Database
	q  Queue
	c  Cache
}

func New(db Database, b Queue, c Cache) Repository {
	return Repository{db, b, c}
}

func (r Repository) CreateTransaction(transaction *domain.Transaction) (uint, error) {

	entity := mapper.ToEntity(transaction)

	if err := r.db.Create(&entity).Error; err != nil {
		return 0, err
	}

	r.c.Set(context.Background(), fmt.Sprint(transaction.ID), fmt.Sprint(transaction.Status), time.Minute*1)

	return entity.ID, nil
}

func (r Repository) ExistsByUserIDAndStatus(payerID uint, status []domain.Status) bool {
	err := r.db.Where("payer_id = ? AND status IN ?", payerID, status).First(&domain.Transaction{}).Error
	return err == nil
}

func (r Repository) SendMessage(message *domain.TransactionQueueMessage) error {
	body, _ := json.Marshal(message)
	return r.q.SendMessage(string(body))
}

func (r Repository) GetTransactionByIDAndPayerID(transactionID, payerID uint) (*model.Transaction, error) {
	transaction := new(model.Transaction)
	if err := r.db.Where("id = ? AND payer_id = ?", transactionID, payerID).First(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}
