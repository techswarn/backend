package services

import (
	"github.com/techswarn/backend/models"
    "github.com/google/uuid"
    "time"
   "errors"
	"fmt"
    "github.com/techswarn/backend/database"
	"github.com/techswarn/backend/utils"
)

func CreateItem(ItemRequest models.ItemRequest) (models.Item, error) {
	var item models.Item
	path, err := utils.UploadImage("mackerel.jpg")

	if err != nil {
		fmt.Printf("Error while uploading image to spaces: %s", err)
		return item, err
	}
	fmt.Println("IMAGE PATH: %s \n", path)
	item = models.Item{
		ID: uuid.New().String(),
		Name: ItemRequest.Name,
		Price: ItemRequest.Price,
		Quantity: ItemRequest.Quantity,
		Image: path,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	fmt.Printf("%#v", item)
	
	if result := database.DB.Create(&item); result.Error != nil {
		fmt.Printf("DB write error: %s", &result.Error)
		return item, result.Error
	}

	return item, nil
}

func GetAllItem() []models.Item {
	// create a variable to store items data
	var Item []models.Item = []models.Item{}

	// get all data from the database order by created_at
	database.DB.Order("created_at desc").Find(&Item)

	// return all items from the database
	return Item
}

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
	
		// update the fish data in the database
		database.DB.Save(&item)
	
		// return the updated item
		return item, nil
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