package services

import (
	"fmt"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/database"
	"gorm.io/gorm"

)

func AddInventory(itemRequest *models.Inventory_request) (string, error) {
	fmt.Printf("checking input item to inventory : #%v", itemRequest)
	var item models.ItemUser;
	var result *gorm.DB
	for _, v := range itemRequest.ItemID {
		item = models.ItemUser{
			UserID : itemRequest.UserID,
			ItemID : v,
		}

		result = database.DB.Create(&item); 
		
		if result.Error != nil {
			fmt.Printf("DB write error: %s", &result.Error)
			return "failed", result.Error
		}
	}




	return "successfull", nil
}