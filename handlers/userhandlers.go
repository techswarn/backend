package handlers

import(
	"net/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/services"
	"github.com/techswarn/backend/models"
	_ "github.com/techswarn/backend/utils"
)

func GetUser(c *fiber.Ctx) error {
	var userId = c.Params("id")
	user, err := services.GetUserByID(userId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	fmt.Printf("%#v", user)
	return c.JSON(models.Response[any]{
		Success: true,
		Message: "User for the id",
		Data:    user,
	})
}

func GetAllSellersByType(c *fiber.Ctx) error {
	sellers := services.GetAllSellersByType("sellers")

	fmt.Printf("List of Sellers: %#v \n", sellers)
	

	return c.Status(http.StatusOK).JSON(models.Response[[]models.User]{
		Success: true,
		Message: "List of item",
		Data:  sellers,
	})
}