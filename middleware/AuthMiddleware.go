package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleware "github.com/gofiber/jwt/v3"
	"github.com/techswarn/backend/utils"
)

// CreateMiddleware returns a middleware with JWT authentication
func CreateMiddleware() func(*fiber.Ctx) error {
	
	// create a JWT middleware
	config := jwtMiddleware.Config{
		SigningKey:   []byte(utils.GetValue("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	// return the JWT middleware
	return jwtMiddleware.New(config)
}

// jwtError returns error handler for JWT middleware
func jwtError(c *fiber.Ctx, err error) error {
	// if the error is caused by malformed JWT token
	// return an error
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// if the error is caused by other error
	// return an error
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": err.Error(),
	})
}