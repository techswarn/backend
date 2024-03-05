package handlers

import (
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/utils"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
)

func CreateNewItem(c *fiber.Ctx) error {
	fmt.Println("Create new Item handler")

	//check the token
	isValid, err := utils.CheckToken(c)

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	//Create a variable for Item from request

	var itemInput *models.ItemRequest = new(models.ItemRequest)
	fmt.Printf("checking input item Object: %v \n", itemInput)
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
	 createdItem, err := services.CreateItem(*itemInput)

	if( err != nil ) {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Create Item failed",
		})
	}
	
	// return the created item in a storage
	return c.Status(http.StatusCreated).JSON(models.Response[models.Item]{
		Success: true,
		Message: "item created",
		Data:    createdItem,
	})

}

func GetAllItem(c *fiber.Ctx) error {
	var item []models.Item = services.GetAllItem()

	fmt.Printf("List of items: %#v \n", item)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "List of item",
		Data:  item,
	})
}

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
