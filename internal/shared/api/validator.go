package api

import "github.com/go-playground/validator/v10"

type ValidatorErrorResponse struct {
	Field        string `json:"field"`
	Tag          string `json:"tag"`
	Value        string `json:"value"`
	ErrorMessage string `json:"message"`
}

func ValidateRequest(req interface{}) []*ValidatorErrorResponse {
	var errors []*ValidatorErrorResponse

	var validate = validator.New()
	if err := validate.Struct(req); err != nil {

		for _, validationError := range err.(validator.ValidationErrors) {
			var errorResponse ValidatorErrorResponse
			errorResponse.Field = validationError.Field()
			errorResponse.Tag = validationError.Tag()
			errorResponse.Value = validationError.Param()
			errorResponse.ErrorMessage = validationError.Error()

			errors = append(errors, &errorResponse)
		}
	}

	return errors
}
