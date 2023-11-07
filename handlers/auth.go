package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
)

// Signup return JWT token
func Signup(c *fiber.Ctx) error {
    // create a variable to store the request
	var userInput *models.UserRequest = new(models.UserRequest)

    // parse the request into "userInput" variable
	if err := c.BodyParser(userInput); err != nil {
        // if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // validate the request
	errors := userInput.ValidateStruct()

    // if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	fmt.Printf("%#v", userInput)

    if userInput.Password != userInput.ConfirmPassword {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Password does not match",
		})
	}

    // perform the signup
    // if signup is successful, the JWT token is returned
	token, err := services.Signup(*userInput)

    // if signup is failed, return an error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // return the JWT token
	return c.JSON(models.Response[string]{
		Success: true,
		Message: "token data",
		Data:    token,
	})
}

// Login returns JWT token for registered user
func Login(c *fiber.Ctx) error {
	fmt.Println("Here 1")
    // create a variable to store the request
	var userInput *models.UserLoginRequest = new(models.UserLoginRequest)

    // parse the request into "userInput" variable
	if err := c.BodyParser(userInput); err != nil {
         // if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // validate the request
	errors := userInput.ValidateStruct()
    fmt.Printf("%#v", userInput)
    // if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

    // perform the login
    // if login is successful, the JWT token is returned
	token, err := services.Login(*userInput)

    // if login is failed, return an error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

    // return the JWT token
	return c.JSON(models.Response[string]{
		Success: true,
		Message: "token data",
		Data:    token,
	})
}
