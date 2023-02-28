package domain

type TransactionMessage struct {
	Value         float64
	Payer         uint
	Payee         uint
	TransactionID uint
}
