package models

type Inventory_request struct {
	UserID string `json:"userid" validate:"required"`
	ItemID []string `json:"itemid" validate:"required"`
}
