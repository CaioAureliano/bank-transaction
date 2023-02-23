package user

type CreateRequestDTO struct {
	Firstname string `json:"firstname" validate:"required,min=3"`
	Lastname  string `json:"lastname" validate:"required,min=3"`
	Email     string `json:"email" validate:"required,email"`
	CPF       string `json:"cpf" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Type      string `json:"type" validate:"required"`
}
