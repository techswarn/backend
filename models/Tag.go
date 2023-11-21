package models

import "github.com/go-playground/validator/v10"

type Tag struct {
	TagID string `json:"tagid"`
	Name string `json:"name"`
}

type TagRequest struct {
	Name string `json:"name" validate:"required"`
}

//ValidateStruct returns validation errors if validation failed
func (TagInput TagRequest) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(TagInput)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}

	return errors
}