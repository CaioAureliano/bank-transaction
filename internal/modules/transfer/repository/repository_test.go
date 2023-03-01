package repository

import (
	"fmt"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/transfer/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/CaioAureliano/bank-transaction/pkg/utils/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupBefore() (*gorm.DB, sqlmock.Sqlmock) {
	conn, mock, _ := sqlmock.New()
	db := database.Connection(test.DialectorMock(conn))
	return db, mock
}

func TestUpdateAccounts(t *testing.T) {
	t.Parallel()

	const queryUpdate = "UPDATE `accounts` SET (.+) WHERE `id` = ?"

	t.Run("should be commit", func(t *testing.T) {
		db, mock := setupBefore()
		r := New(db, nil)

		transference := &domain.Transference{
			Payer: &model.Account{ID: 1},
			Payee: &model.Account{ID: 2},
		}

		mock.ExpectBegin()
		mock.ExpectExec(queryUpdate).
			WithArgs(transference.Payer.UserID, transference.Payer.Type, transference.Payer.Balance, transference.Payer.ID).
			WillReturnResult(sqlmock.NewResult(int64(transference.Payer.ID), 1))

		mock.ExpectExec(queryUpdate).
			WithArgs(transference.Payee.UserID, transference.Payee.Type, transference.Payee.Balance, transference.Payee.ID).
			WillReturnResult(sqlmock.NewResult(int64(transference.Payer.ID), 1))
		mock.ExpectCommit()

		err := r.UpdateAccounts(transference)

		assert.NoError(t, err)
	})

	t.Run("should be rollback with error", func(t *testing.T) {
		db, mock := setupBefore()
		r := New(db, nil)

		transference := &domain.Transference{
			Payer: &model.Account{ID: 1},
			Payee: &model.Account{ID: 2},
		}

		mock.ExpectBegin()
		mock.ExpectExec(queryUpdate).
			WithArgs(transference.Payer.UserID, transference.Payer.Type, transference.Payer.Balance, transference.Payer.ID).
			WillReturnResult(sqlmock.NewResult(int64(transference.Payer.ID), 1))

		mock.ExpectExec(queryUpdate).
			WithArgs(transference.Payee.UserID, transference.Payee.Type, transference.Payee.Balance, transference.Payee.ID).
			WillReturnError(fmt.Errorf("ops: mock error"))
		mock.ExpectRollback()

		err := r.UpdateAccounts(transference)

		assert.Error(t, err)
	})
}

func TestGetAccountByID(t *testing.T) {
	t.Parallel()

	db, mock := setupBefore()
	r := New(db, nil)

	userIDmock := uint(1)
	expectedAccount := &model.Account{
		ID:      1000,
		UserID:  1,
		Type:    model.USER,
		Balance: 1000,
	}

	columns := []string{"Account__id", "Account__user_id", "Account__type", "Account__balance"}
	rows := sqlmock.NewRows(columns).AddRow(expectedAccount.ID, expectedAccount.UserID, expectedAccount.Type, expectedAccount.Balance)

	mock.ExpectQuery("SELECT (.*) FROM `users` LEFT JOIN `accounts` (.+) WHERE `users`.`id` = ? (.+)").
		WithArgs(userIDmock).
		WillReturnRows(rows)

	account, err := r.GetAccountByID(userIDmock)

	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, expectedAccount, account)
}
