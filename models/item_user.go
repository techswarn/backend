package models

type ItemUser struct {
	ID     int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`  
	UserID string `json:"userid" gorm:"foreignKey"`
	ItemID string `json:"itemid" gorm:"foreignKey"`
}