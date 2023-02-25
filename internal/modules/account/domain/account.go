package domain

type Account struct {
	ID uint
	*User
	Type
	Balance float64
}
