package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Database interface {
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Save(value interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
}

type Cache interface {
	Set(ctx context.Context, k string, v interface{}, expiration time.Duration) *redis.StatusCmd
}

type Repository struct {
	db Database
	c  Cache
}

func New(db Database, c Cache) Repository {
	return Repository{db, c}
}

func (r Repository) GetAccountByID(userID uint) (*model.Account, error) {
	user := new(model.User)
	if result := r.db.Joins("Account").First(&user, model.User{ID: userID}); result.Error != nil || user == nil {
		return nil, errors.New(fmt.Sprintf("not found user by userID(%d): %s", userID, result.Error.Error()))
	}
	return user.Account, nil
}

func (r Repository) CacheStatus(status model.Status, transactionID uint) error {
	return r.c.Set(context.Background(), fmt.Sprint(transactionID), fmt.Sprint(status), time.Minute*1).Err()
}

func (r Repository) UpdateAccounts(transference *domain.Transference) error {
	return r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Save(&transference.Payer).Error; err != nil {
			return err
		}

		if err := tx.Save(&transference.Payee).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r Repository) UpdateTransaction(transaction *model.Transaction) error {
	return r.db.Model(&model.Transaction{ID: transaction.ID}).Updates(transaction).Error
}

func (r Repository) Authenticator() error {
	return nil
}