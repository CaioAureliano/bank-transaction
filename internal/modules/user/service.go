package user

type Service interface {
	Create(CreateRequestDTO) error
}

type service struct{}

func NewService() Service {
	return service{}
}

func (s service) Create(userDTO CreateRequestDTO) error {

	return nil
}
