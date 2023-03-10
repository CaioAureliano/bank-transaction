package dto

type TransactionRequestDTO struct {
	Value float64 `json:"value" validate:"required"`
	Payee uint    `json:"payee" validate:"required"`
}

type GetTransactionRequestDTO struct {
	TransactionID string
	PayerID       uint
}
