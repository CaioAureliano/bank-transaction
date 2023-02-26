package service

import (
	"log"
	"time"

	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/dto"
	"github.com/CaioAureliano/bank-transaction/internal/modules/account/domain/mapper"
	"github.com/CaioAureliano/bank-transaction/pkg/authentication"
)

type Service struct {
	r repository
	v validator
}

func New(r repository, v validator) Service {
	return Service{r, v}
}

func (s Service) CreateUserAccount(req dto.CreateRequestDTO) error {

	user := mapper.RequestToModel(req)

	if err := s.v.Validate(user); err != nil {
		log.Printf("error to validate user - %s", err)
		return err
	}

	user.GeneratePassword()

	if err := s.r.Create(user); err != nil {
		log.Printf("error to try create user - %s", err)
		return err
	}

	return nil
}

func (s Service) Authenticate(req dto.AuthRequestDTO) (string, error) {

	entity, err := s.r.GetByEmail(req.Email)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	user := mapper.ToModel(entity)
	if err := user.ValidatePassword(req.Password); err != nil {
		log.Printf("invalid password - %s", err)
		return "", err
	}

	return authentication.GenerateJwt(user.ID, uint(user.Account.Type), time.Now().Add(time.Hour*12))
}
