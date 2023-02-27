package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/database"
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

	email := "example@mail.com"
	rowsMock := sqlmock.NewRows([]string{"email"}).AddRow(email)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT `users`.`id`,`users`.`firstname`,`users`.`lastname`,`users`.`cpf`,`users`.`email`,`users`.`password`,`users`.`created_at`,`Account`.`id` AS `Account__id`,`Account`.`user_id` AS `Account__user_id`,`Account`.`type` AS `Account__type`,`Account`.`balance` AS `Account__balance` FROM `users` LEFT JOIN `accounts` `Account` ON `users`.`id` = `Account`.`user_id` WHERE email = ? ORDER BY `users`.`id` LIMIT 1")).
		WithArgs(email).
		WillReturnRows(rowsMock)

	user, err := r.GetByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
}
