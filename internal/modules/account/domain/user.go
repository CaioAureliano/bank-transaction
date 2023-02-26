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

func (u *User) GeneratePassword() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
}
