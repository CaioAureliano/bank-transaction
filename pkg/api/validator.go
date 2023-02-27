package api

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type ValidatorErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type ValidatorErrorsResponse []*ValidatorErrorResponse

func (v *ValidatorErrorsResponse) String() string {
	s, _ := json.Marshal(v)
	return string(s)
}

func ValidateRequest(req interface{}) ValidatorErrorsResponse {
	var errors []*ValidatorErrorResponse

	var validate = validator.New()
	if err := validate.Struct(req); err != nil {

		for _, validationError := range err.(validator.ValidationErrors) {
			var errorResponse ValidatorErrorResponse
			errorResponse.Field = validationError.Field()
			errorResponse.Tag = validationError.Tag()
			errorResponse.Value = validationError.Param()

			errors = append(errors, &errorResponse)
		}
	}

	return errors
}
