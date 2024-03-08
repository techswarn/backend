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
	Phone int64 `JSON:"phone" gorm:"unique"`
	Access string `JSON:"access"`
	Type string `JSON:"type"`
	Password  string    `json:"password"`
	Photo  string `json:"photo"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}