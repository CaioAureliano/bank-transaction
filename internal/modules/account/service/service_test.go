package service

import (
	"errors"
	"testing"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockRepository struct {
	fnCreate func(*domain.User) error
}

func (m mockRepository) Create(user *domain.User) error {
	if m.fnCreate == nil {
		return nil
	}
	return m.fnCreate(user)
}

type mockValidator struct {
	fnValidate func(*domain.User) error
}

func (m mockValidator) Validate(u *domain.User) error {
	if m.fnValidate == nil {
		return nil
	}
	return m.fnValidate(u)
}

func TestCreateUserAccount(t *testing.T) {

	t.Run("Should be mapped reques to domain model", func(t *testing.T) {
		t.Parallel()

		expectedUserAccountMapped := &domain.User{
			Firstname: "User",
			Lastname:  "Test",
			Email:     "example@mail.com",
			CPF:       "000.000.000-00",
			Password:  "test1234",
			Account: &domain.Account{
				Type: domain.USER,
			},
		}

		userMapped := &domain.User{}
		validatorMock := mockValidator{
			fnValidate: func(u *domain.User) error {
				userMapped = u
				return nil
			},
		}

		repositoryMock := mockRepository{}
		reqMock := dto.CreateRequestDTO{
			Firstname: "User",
			Lastname:  "Test",
			Email:     "example@mail.com",
			CPF:       "000.000.000-00",
			Password:  "test1234",
			Type:      0,
		}

		s := New(repositoryMock, validatorMock)
		err := s.CreateUserAccount(reqMock)

		assert.NoError(t, err)
		assert.Equal(t, expectedUserAccountMapped.Firstname, userMapped.Firstname)
		assert.Equal(t, expectedUserAccountMapped.Lastname, userMapped.Lastname)
		assert.Equal(t, expectedUserAccountMapped.CPF, userMapped.CPF)
		assert.Equal(t, expectedUserAccountMapped.Email, userMapped.Email)
		assert.Equal(t, expectedUserAccountMapped.Account.Type, userMapped.Account.Type)
	})

	t.Run("Should be correct generate password", func(t *testing.T) {
		t.Parallel()

		passwordMock := "test1234"

		var hashPasswordGenerated string
		validatorMock := mockValidator{}
		repositoryMock := mockRepository{func(u *domain.User) error {
			hashPasswordGenerated = u.Password
			return nil
		}}
		reqMock := dto.CreateRequestDTO{
			Password: passwordMock,
		}

		s := New(repositoryMock, validatorMock)
		err := s.CreateUserAccount(reqMock)

		assert.NoError(t, err)
		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(hashPasswordGenerated), []byte(passwordMock)))
	})

	t.Run("Should be correct errors return", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name string

			validatorMock  validator
			repositoryMock repository
			reqMock        dto.CreateRequestDTO

			expectedError assert.ErrorAssertionFunc
		}{
			{
				name: "should be create user account with valid request dto",

				validatorMock: mockValidator{},
				repositoryMock: mockRepository{
					fnCreate: func(u *domain.User) error {
						return nil
					},
				},
				reqMock: dto.CreateRequestDTO{},

				expectedError: assert.NoError,
			},
			{
				name: "should be return validation error",

				validatorMock: mockValidator{
					fnValidate: func(u *domain.User) error {
						return errors.New("invalid")
					},
				},
				repositoryMock: mockRepository{
					fnCreate: func(u *domain.User) error {
						return nil
					},
				},
				reqMock: dto.CreateRequestDTO{},

				expectedError: assert.Error,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				s := New(tt.repositoryMock, tt.validatorMock)
				err := s.CreateUserAccount(tt.reqMock)

				tt.expectedError(t, err)
			})
		}
	})
}
