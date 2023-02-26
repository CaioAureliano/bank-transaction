package model

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Firstname string
	Lastname  string
	CPF       string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	Account   *Account
	CreatedAt time.Time
}
