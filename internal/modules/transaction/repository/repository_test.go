package repository

import (
	"regexp"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/utils/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	conn, mock, _ := sqlmock.New()
	defer conn.Close()
	db := database.Connection(test.DialectorMock(conn))
	r := New(db, nil)

	transactionToPersist := &domain.Transaction{
		PayerID: 1,
		PayeeID: 2,
		Value:   1000,
		Status:  domain.REQUESTED,
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `transactions` (`payer_id`,`payee_id`,`value`,`status`,`created_at`) VALUES (?,?,?,?,?)")).
		WithArgs(transactionToPersist.PayerID, transactionToPersist.PayeeID, transactionToPersist.Value, transactionToPersist.Status, test.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(100, 1))
	mock.ExpectCommit()

	transactionID, err := r.CreateTransaction(transactionToPersist)

	assert.NoError(t, err)
	assert.Equal(t, uint(100), transactionID)
}
