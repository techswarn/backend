package models

import (
	"github.com/go-playground/validator/v10"
	"fmt"
)

type ItemRequest struct {
	Name     string `json:"name" validate:"required"`
	Price    int    `json:"price" validate:"required,gt=0"`
	Quantity int    `json:"quantity" validate:"gte=0"`
}

// ValidateStruct performs struct based validation
func (itemInput ItemRequest) ValidateStruct() []*ErrorResponse {
	// create a variable to store validation errors
	fmt.Println("Code reached here")
	var errors []*ErrorResponse
	// create a new validator
	validate := validator.New()
	fmt.Println("Code reached here2")
	// validate the struct
	err := validate.Struct(itemInput)
	fmt.Println("Code reached here3")
	// if the validation is failed
    // insert the error inside "errors" variable
	fmt.Printf("%v", err)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}
    // return the validation errors
	return errors
}