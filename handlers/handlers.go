package handlers

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"runtime"
	"time"
	"fmt"
)

func Process(c *fiber.Ctx) error {

	log.Println("Health checks route")
	headers := c.GetReqHeaders()
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {

		go func() {
				for {
					select {
					case <-done:
						log.Printf("CPU count: %d \n", runtime.NumCPU())
						log.Printf("GO routine count: %d \n", runtime.NumGoroutine())
						return
					default:
					}
				}
		}()
	}
	time.Sleep(time.Second * 10)
	close(done)
	return c.Status(http.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "backend api",
		Data: headers,
	})
}

//IMAGE UPLOAD HANDLERS

func UploadFile(c *fiber.Ctx) error {
	  // Get first file from form field "document":
	  file, err := c.FormFile("image")

	  if err!= nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	  }

	  // Save file to root directory:
	  return c.SaveFile(file, fmt.Sprintf("./../assets/%s", file.Filename))
}