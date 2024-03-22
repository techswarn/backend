package handlers

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"runtime"
	"time"
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