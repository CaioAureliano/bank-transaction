package model

import "time"

type Transaction struct {
	ID      uint `gorm:"primaryKey"`
	PayerID uint
	PayeeID uint
	Value   float64
	Status
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
