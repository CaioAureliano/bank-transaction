package repository

import (
	"database/sql"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
	"github.com/CaioAureliano/bank-transaction/pkg/utils/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func before() (Repository, *sql.DB, sqlmock.Sqlmock) {
	conn, mock, _ := sqlmock.New()
	db := database.Connection(test.DialectorMock(conn))
	r := New(db)
	return r, conn, mock
}

func TestCreate(t *testing.T) {
	r, conn, mock := before()
	defer conn.Close()

	user := &domain.User{
		Firstname: "Ada",
		Lastname:  "Lovelace",
		CPF:       "000.000.000-00",
		Email:     "example@mail.com",
		Password:  "test1234",
		Account: &domain.Account{
			Balance: 200,
			Type:    domain.USER,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users`").
		WithArgs(user.Firstname, user.Lastname, user.CPF, user.Email, user.Password, test.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `accounts`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := r.Create(user)

	assert.NoError(t, err)
}

func TestGetByEmail(t *testing.T) {
	r, conn, mock := before()
	defer conn.Close()

	expectedUser := &model.User{
		ID:       1,
		Password: "hash",
		Account: &model.Account{
			Type: model.USER,
		},
	}
	columns := []string{"id", "password", "Account__type"}
	rows := sqlmock.NewRows(columns).AddRow(expectedUser.ID, expectedUser.Password, expectedUser.Account.Type)

	mock.ExpectQuery("SELECT (.+) FROM (.+) WHERE email = ? (.+) LIMIT 1").
		WithArgs(expectedUser.Email).
		WillReturnRows(rows)

	user, err := r.GetByEmail(expectedUser.Email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)
}
