package handlers

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
	"github.com/techswarn/backend/utils"
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

func Checkauth(c *fiber.Ctx) error {
	fmt.Println("Here 1")
	var token string = c.Get("Token")
	_ = token
	isValid, err := utils.CheckToken(c)

	userid, err := utils.GetClaimData(c)

	fmt.Printf("Userid : %s \n", userid)
	fmt.Printf("Is token value: %v \n", isValid)
	user, err := services.GetUserByID(userid)
	if err != nil {
		fmt.Printf("Error while fetching user: %s", err)
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		}) 
	}

	// if token is not valid, return an error
	if !isValid {
		return c.Status(http.StatusUnauthorized).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "Valid user",
		Data: user,
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
	data, err := services.Login(*userInput)

    // if login is failed, return an error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
        Name:     "token",
        Value:    data.Token,
        Expires:  time.Now().Add(24 * time.Hour),
        HTTPOnly: true,
        SameSite: "lax",
    })

	// return the JWT token
	return c.JSON(models.Response[any]{
		Success: true,
		Message: "token data",
		Data:    data,
	})

}
