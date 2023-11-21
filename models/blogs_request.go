package models

import (
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Blog_request struct {
	UserID  string `json:"userid", validate:"required"`
	Subject string `json:"subject", validate:"required"`
	Paragraph string `json:"paragraph", validate:"required"`
}

func (blog_details *Blog_request) ValidateStruct() []*ErrorResponse{

	fmt.Printf("Blog request %#v \n", blog_details)
	var errors []*ErrorResponse
	validate := validator.New()

	// validate the struct
	err := validate.Struct(blog_details)

	// if the validation is failed
    // insert the error inside "errors" variable

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}
    // return the validation errors
	return errors
}