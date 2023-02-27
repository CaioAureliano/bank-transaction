package mapper

import (
	"github.com/CaioAureliano/bank-transaction/internal/modules/transaction/domain"
	"github.com/CaioAureliano/bank-transaction/pkg/model"
)

func ToEntity(t *domain.Transaction) *model.Transaction {
	return &model.Transaction{
		ID:        t.ID,
		PayerID:   t.PayerID,
		PayeeID:   t.PayeeID,
		Value:     t.Value,
		Status:    model.Status(t.Status),
		CreatedAt: t.CreatedAt,
	}
}
