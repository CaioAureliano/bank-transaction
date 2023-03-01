package domain

type TransactionQueueMessage struct {
	Value         float64
	Payer         uint
	Payee         uint
	TransactionID uint
}
