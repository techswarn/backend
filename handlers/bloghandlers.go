package handlers

import (

	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"github.com/techswarn/backend/services"
)

type Search struct {
	Keyword string `json:"keyword"`
	Tag string `json:"tag"`
}

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

func GetBlogs(c *fiber.Ctx) error {
	s := new(Search)
	if err := c.BodyParser(s); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}
	fmt.Println("------------")
	fmt.Println("ontrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.")
	fmt.Println("------------")
	var blog []models.Blog = services.GetBlogs(s.Keyword, s.Tag)
	fmt.Printf("List of Blog: %#v \n", blog)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Blog]{
		Success: true,
		Message: "List of blogs",
		Data:  blog,
	})
}
