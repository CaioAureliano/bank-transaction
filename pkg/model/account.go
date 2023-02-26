package model

type Account struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Type
	Balance float64
}
