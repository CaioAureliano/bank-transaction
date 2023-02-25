package handler

import "github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"

type mockService struct {
	fnCreate func(dto.CreateRequestDTO) error
}

func (m mockService) Create(req dto.CreateRequestDTO) error {
	if m.fnCreate == nil {
		return nil
	}
	return m.fnCreate(req)
}
