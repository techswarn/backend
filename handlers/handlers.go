package handlers

import (
	"fmt"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/backend/models"
	"runtime"

	"sync"
)

func Process(c *fiber.Ctx) error {
	wg := sync.WaitGroup{}
	fmt.Println("Health checks route")
	fmt.Printf("CPU count: %d \n", runtime.NumCPU())
	fmt.Printf("GO routine count: %d \n", runtime.NumGoroutine())
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Printf("CPU count: %d \n", runtime.NumCPU())
					fmt.Printf("GO routine count: %d \n", runtime.NumGoroutine())
					return
				default:
				}
				fmt.Println(n)
			}
		}(i)
	}
	wg.Wait()
//	time.Sleep(time.Second * 10)
	close(done)
	return c.Status(http.StatusOK).JSON(models.Response[any]{
		Success: true,
		Message: "backend api",
	})
}