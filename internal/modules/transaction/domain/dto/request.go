package dto

type TransactionRequestDTO struct {
	Value float64 `json:"value" validate:"required"`
	Payer uint    `json:"payer" validate:"required"`
	Payee uint    `json:"payee" validate:"required"`
}
