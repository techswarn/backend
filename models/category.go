package models

type Category struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Category  string   `json:"category"`
}