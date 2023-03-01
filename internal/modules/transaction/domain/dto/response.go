package dto

import "github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"

type TransactionResponseDTO struct {
	Status  domain.Status `json:"status"`
	Message string        `json:"message"`
}
