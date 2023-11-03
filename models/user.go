package models

import (
	"time"
)

type User struct {
	ID string `JSON:"id"`
	FirstName string `JSON:"fistname"`
	LastName string `JSON:"lastname"`
	UserName string `JSON:"username"`
	Email string `JSON:"email" gorm:"unique"`
	Access string `JSON:"access"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}