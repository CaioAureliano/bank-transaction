package handler

import "github.com/CaioAureliano/bank-transaction/internal/modules/user"

type mockService struct {
	fnCreate func(user.CreateRequestDTO) error
}

func (m mockService) Create(dto user.CreateRequestDTO) error {
	if m.fnCreate == nil {
		return nil
	}
	return m.fnCreate(dto)
}
