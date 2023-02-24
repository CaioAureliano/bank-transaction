package service

import "github.com/CaioAureliano/bank-transaction/internal/modules/user/domain/dto"

type Service interface {
	Create(dto.CreateRequestDTO) error
}

type userService struct{}

func NewService() Service {
	return userService{}
}

func (s userService) Create(userDTO dto.CreateRequestDTO) error {

	return nil
}
