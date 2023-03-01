package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint
	Firstname string
	Lastname  string
	CPF       string
	Email     string
	Password  string
	Account   *Account
	CreatedAt time.Time
}

func (u *User) GenerateHash() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
}

func (u *User) ValidatePassword(p string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
}
