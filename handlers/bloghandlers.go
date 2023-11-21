package handlers

import (

	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
)

func CreateNewBlog (c *fiber.Ctx) error {
	fmt.Println("Create new Blog")

	//Create a request from blog model
	var blogInput *models.Blog_request = new(models.Blog_request)

	if err := c.BodyParser(blogInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	fmt.Printf("Blog input %#v", blogInput)
	// validate the request
	errors := blogInput.ValidateStruct()
	fmt.Printf("validate struct error: %#v ",errors )

	if (errors != nil) {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	var b, err = services.CreateBlog(*blogInput)

	if( err != nil ) {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Create Blog failed",
		})
	}

	fmt.Printf("Blog Created response %#v \n", b)


	
	return c.Status(http.StatusCreated).JSON(models.Response[models.Blog]{
		Success: true,
		Message: "item created",
		Data:    b,
	})
}

func GetAllBlogs(c *fiber.Ctx) error {
	var blogs []models.Blog = services.GetAllBlogs()
	fmt.Printf("List of Blogs: %#v \n", blogs)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Blog]{
		Success: true,
		Message: "List of blogs",
		Data:  blogs,
	})
}