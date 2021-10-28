package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	counter := 0

	app.Get("/api", func(c *fiber.Ctx) error {
		counter++

		if counter >= 5 && counter <= 10 {
			time.Sleep(time.Millisecond * 1000)
		}

		msg := fmt.Sprintf("Hello %v", counter)
		fmt.Println(msg)

		return c.SendString(msg)
	})

	app.Listen(":8000")
}
