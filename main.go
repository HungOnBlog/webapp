package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/long", func(c *fiber.Ctx) error {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
