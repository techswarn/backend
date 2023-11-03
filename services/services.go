package services

import (
	"github.com/techswarn/backend/models"
    "github.com/google/uuid"
    "time"
    "errors"
	"fmt"
    "github.com/techswarn/backend/database"
)

var storage []models.Item = []models.Item{}

// GetAllItems returns all items data
func GetAllItems() []models.Item {
	// create a variable to store items data
	var items []models.Item = []models.Item{}

	// get all data from the database order by created_at
	database.DB.Order("created_at desc").Find(&items)

	// return all items from the database
	return items
}

// GetItemByID returns get item's data by ID
func GetItemByID(id string) (models.Item, error) {
    // create a variable to store item data
	var item models.Item

    	// get item data from the database by ID
	result := database.DB.First(&item, "id = ?", id)

	// if the item data is not found, return an error
	if result.RowsAffected == 0 {
		return models.Item{}, errors.New("item not found")
	}

	// return the item data from the database
	return item, nil
}

//Create a new item
func CreateItem(itemRequest models.ItemRequest) models.Item {
    var newItem models.Item = models.Item{
        ID: uuid.New().String(),
        Name: itemRequest.Name,    
        Price:  itemRequest.Price,  
        Quantity:  itemRequest.Quantity,
        CreatedAt: time.Now(),
    }
    fmt.Println(newItem)

	// insert the new item data into the database
	database.DB.Create(&newItem)

	// return the recently inserted item
	return newItem
}

// UpdateItem returns updated item
func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	// get the item data by ID
	item, err := GetItemByID(id)

	// if item is not found, return an error
	if err != nil {
		return models.Item{}, err
	}

	// update item data
	item.Name = itemRequest.Name
	item.Price = itemRequest.Price
	item.Quantity = itemRequest.Quantity
	item.UpdatedAt = time.Now()

	// update the item data in the database
	database.DB.Save(&item)

	// return the updated item
	return item, nil
}

// DeleteItem returns deletion result
func DeleteItem(id string) bool {
	// get the item data by ID
	item, err := GetItemByID(id)

	// if item is not found, return false
	if err != nil {
		return false
	}

	// delete the item data
	database.DB.Delete(&item)

	// return true
	// this means the deletion is succeed
	return true
}