package models

import "time"

type Blog struct {
	ID 					string `json:"id"`
	UserID 				string `json:"userid"`
	Subject 			string `json:"subject"`
	Paragraph 			string `json:"paragraph"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
