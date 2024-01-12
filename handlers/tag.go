package handlers

import (
	"fmt"
	"github.com/techswarn/backend/models"
	"net/http"
	"github.com/techswarn/backend/services"
	"github.com/gofiber/fiber/v2"	
)

func CreateNewTag(c *fiber.Ctx) error {
	fmt.Println("Handler: TAG | CreateNewTag")
	
	//Create a request from Model
	var tagInput *models.Tag_Request = new(models.Tag_Request)

	if err := c.BodyParser(tagInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	fmt.Printf("Tag request body: %#v", tagInput)

	errors := tagInput.ValidateStruct()

	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	// Write a method in services and invoke here.
		t, err := services.CreateTag(*tagInput)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.Response[[]*models.ErrorResponse]{
				Success: false,
				Message: "Create tag failed",
			})
		}

	return c.Status(http.StatusCreated).JSON(models.Response[models.Tag]{
		Success: true,
		Message: "tag created",
		Data:    t,
	})
}