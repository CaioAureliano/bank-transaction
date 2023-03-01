package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain/dto"
	"github.com/CaioAureliano/bank-transaction/pkg/configuration"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Database interface {
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Joins(query string, args ...interface{}) *gorm.DB
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
	Save(value interface{}) *gorm.DB
	Updates(values interface{}) *gorm.DB
	Model(value interface{}) *gorm.DB
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

var (
	ErrFailedAuthenticator = errors.New("failed to on authenticator")
)

func (r Repository) GetAccountByID(userID uint) (*model.Account, error) {
	user := new(model.User)
	if err := r.db.Joins("Account").First(&user, model.User{ID: userID}).Error; err != nil || user == nil {
		return nil, errors.New(fmt.Sprintf("not found user with id \"%d\": %s", userID, err.Error()))
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

	c := fiber.AcquireClient()

	agent := c.Get(configuration.Env.AUTHENTICATORSERVICEURL)
	if err := agent.Parse(); err != nil {
		log.Println(err)
		return ErrFailedAuthenticator
	}

	res := new(dto.AuthenticatorResponseDTO)

	code, _, errs := agent.Struct(&res)
	if !res.Authorization || code != 200 || len(errs) > 0 {
		log.Printf("errors: %v", errs)
		return ErrFailedAuthenticator
	}

	return nil
}
