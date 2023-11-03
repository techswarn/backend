package handlers

import (
    "net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/services"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/utils"

	"fmt"
)

// GetAllItems returns all items from the storage
func GetAllItems(c *fiber.Ctx) error {
    // get all items
	var items []models.Item = services.GetAllItems()

    // return the response
	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items data",
		Data:    items,
	})
}

//GetAllItems by ID

func GetItemByID(c *fiber.Ctx) error {
	var itemId = c.Params("id")
	item , err := services.GetItemByID(itemId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item for the id",
		Data:    item,
	})
}

//Create new Item 

func CreateNewItem(c *fiber.Ctx) error {
	fmt.Println("Create new item handler")

	// check the token
	isValid, err := utils.CheckToken(c)

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	//Create a variable for item from request

	var itemInput *models.ItemRequest = new(models.ItemRequest)
	fmt.Printf("%v", itemInput)
	// parse the request into "itemInput" variable
	if err := c.BodyParser(itemInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	// validate the request
	errors := itemInput.ValidateStruct()

	// if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	// create a new item
	var createdItem models.Item = services.CreateItem(*itemInput)

	// return the created item in a storage
	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})
}

// UpdateItem returns updated item
func UpdateItem(c *fiber.Ctx) error {

	// check the token
	isValid, err := utils.CheckToken(c)

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // create a variable to store the request
	var itemInput *models.ItemRequest = new(models.ItemRequest)

    // parse the request into "itemInput" variable
	if err := c.BodyParser(itemInput); err != nil {
        // if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // validate the request
    errors := itemInput.ValidateStruct()

    // if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

    // get the item's ID from the request parameter
	var itemID string = c.Params("id")

    // update the item's data
	updatedItem, err := services.UpdateItem(*itemInput, itemID)
	if err != nil {
        // if update is failed, return an error
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // return the updated item
	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item updated",
		Data:    updatedItem,
	})
}

// DeleteItem returns deletion result
func DeleteItem(c *fiber.Ctx) error {

	// // check the token
	isValid, err := utils.CheckToken(c)

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // get the item's ID from the request parameter
	var itemID string = c.Params("id")

    // delete the item data
	var result bool = services.DeleteItem(itemID)

	if result {
        // if successful, return the result
		return c.JSON(models.Response[any]{
			Success: true,
			Message: "item deleted",
		})
	}

    // return the failed result
	return c.Status(http.StatusNotFound).JSON(models.Response[any]{
		Success: false,
		Message: "item failed to delete",
	})
}