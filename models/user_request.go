package models

import "github.com/go-playground/validator/v10"

type UserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	FirstName string `json:"firstname" validate:"required"`
	LastName string `json:"lastname" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
    Access string `json:"access"`
	Phone int64 `JSON:"phone" validate:"required,min=10"`
	Type string `JSON:"type"`
	ConfirmPassword string `json:confirmpassword" validate:"required,min=6"`
	State string `json:"state"`
	City string `json:"city"`
	Street string `json:"street"`
	Pincode int `json:"pincode"`
}

type UserLoginRequest struct {
	Phone 	int64 `JSON:"phone" gorm:"unique"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

//ValidateStruct returns validation errors if validation failed
func (userInput UserRequest) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userInput)
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

//ValidateStruct returns validation errors if validation failed
func (userInput UserLoginRequest) ValidateStruct() []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(userInput)
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
