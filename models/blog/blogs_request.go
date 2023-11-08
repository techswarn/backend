package models

import (
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Blog_request struct {
	Subject string `json:"subject", validate:"required"`
	Paragraph string `json:"paragraph", validate:"required"`
	Tags []string `json:"tags", validate:"required"`
}