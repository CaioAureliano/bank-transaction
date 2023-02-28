package domain

import "time"

type Transaction struct {
	ID      uint
	PayerID uint
	PayeeID uint
	Value   float64
	Status
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
