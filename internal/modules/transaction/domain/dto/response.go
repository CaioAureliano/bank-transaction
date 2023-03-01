package dto

import "github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"

type TransactionResponseDTO struct {
	Status  domain.Status `json:"status"`
	Message string        `json:"message"`
}

type CreatedTransactionResponseDTO struct {
	Message string       `json:"message"`
	Links   LinksHateoas `json:"links"`
}

type LinksHateoas struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}
