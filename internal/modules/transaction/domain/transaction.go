package domain

import "time"

type Transaction struct {
	ID      uint
	PayerID uint
	PayeeID uint
	Value   float64
	Status
	CreatedAt time.Time
}
