package models

type Address struct {
    Address_id int `json:"address_id" gorm:"primaryKey;autoIncrement:true;unique"`
    State string `json:"state"`
    City string `json:"city"`
    Street string `json:"street"`
    Pincode int `json:"pincode"`
    User_id string `json:"user_id"`
}