package repository

import (
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
	"github.com/CaioAureliano/bank-transaction/pkg/utils/test"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	conn, mock, err := sqlmock.New()
	defer conn.Close()
	db := database.Connection(test.DialectorMock(conn))
	r := New(db)

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

	err = r.Create(user)

	assert.NoError(t, err)
}
