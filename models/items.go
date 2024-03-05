package models

import "time"

type Item struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	Image     string    `json:"image"`
	CategoryId	int     `json:"category_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}