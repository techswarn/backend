package handlers

import (
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/utils"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
)

func AddInventory(c *fiber.Ctx) error {

 fmt.Println("Add item to inventory")

 //check the token
 isValid, err := utils.CheckToken(c)

 // if token is not valid, return an error
 if !isValid {
	 return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
		 Success: false,
		 Message: err.Error(),
	 })
 }

  var input *models.Inventory_request = new(models.Inventory_request)


  // parse the request into "itemInput" variable
  if err := c.BodyParser(input); err != nil {
	  // if parsing is failed, return an error
	  return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
		  Success: false,
		  Message: err.Error(),
	  })
  }
  fmt.Printf("input: %v \n", input)
  addItem, err := services.AddInventory(input)

  if err != nil {
	return c.Status(http.StatusCreated).JSON(models.Response[any]{
		Success: false,
		Message: "Failed to add to inventory",
   })
  }
 
   return c.Status(http.StatusCreated).JSON(models.Response[any]{
		Success: true,
		Message: addItem,
   })

}