package model

import "time"

type Transaction struct {
	ID      uint `gorm:"primaryKey"`
	PayerID uint
	PayeeID uint
	Value   float64
	Status
	CreatedAt time.Time
}
