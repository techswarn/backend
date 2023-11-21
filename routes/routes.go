package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/handlers"
	"github.com/techswarn/backend/middleware"
)

// SetupRoutes register routes based on functionalities
func SetupRoutes(app *fiber.App) {

	// public routes
	var publicRoutes fiber.Router = app.Group("/api/v1")

	publicRoutes.Post("/signup", handlers.Signup)
	publicRoutes.Post("/login", handlers.Login)
	publicRoutes.Get("/items", handlers.GetAllItems)
	publicRoutes.Get("/items/:id", handlers.GetItemByID)

	//Blog routes
	publicRoutes.Post("/blogs", handlers.CreateNewBlog)
	publicRoutes.Get("/blogs", handlers.GetAllBlogs)

	//Tag routes
	publicRoutes.Post("/tags", handlers.CreateNewTag)
	publicRoutes.Get("/tags", handlers.GetAllTags)

	// private routes, authentication is required
	var privateRoutes fiber.Router = app.Group("/api/v1", middlewares.CreateMiddleware())

	privateRoutes.Post("/items", handlers.CreateNewItem)
	privateRoutes.Put("/items/:id", handlers.UpdateItem)
	privateRoutes.Delete("/items/:id", handlers.DeleteItem)

}