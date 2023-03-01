package repository

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/utils/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

type mockCache struct {
	fnSet func(ctx context.Context, k string, v interface{}, expiration time.Duration) *redis.StatusCmd
}

func (m mockCache) Set(ctx context.Context, k string, v interface{}, expiration time.Duration) *redis.StatusCmd {
	if m.fnSet == nil {
		return nil
	}
	return m.fnSet(ctx, k, v, expiration)
}

func (m mockCache) Get(ctx context.Context, key string) *redis.StringCmd {
	return nil
}

func TestCreateTransaction(t *testing.T) {
	conn, mock, _ := sqlmock.New()
	db := database.Connection(test.DialectorMock(conn))

	cacheMock := mockCache{
		fnSet: func(ctx context.Context, k string, v interface{}, expiration time.Duration) *redis.StatusCmd {
			return nil
		},
	}

	r := New(db, nil, cacheMock)

	transactionToPersist := &domain.Transaction{
		PayerID: 1,
		PayeeID: 2,
		Value:   1000,
		Status:  domain.REQUESTED,
		Message: "",
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `transactions`")).
		WithArgs(transactionToPersist.PayerID, transactionToPersist.PayeeID, transactionToPersist.Value, transactionToPersist.Status, transactionToPersist.Message, test.AnyTime{}, test.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(100, 1))
	mock.ExpectCommit()

	transactionID, err := r.CreateTransaction(transactionToPersist)

	assert.NoError(t, err)
	assert.Equal(t, uint(100), transactionID)
}
