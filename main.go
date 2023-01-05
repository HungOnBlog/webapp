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
		return c.JSON(fiber.Map{
			"path":    "/",
			"message": "OK",
		})
	})

	app.Get("/long", func(c *fiber.Ctx) error {
		for i := 0; i < 10000; i++ {
			fmt.Println(i)
		}
		return c.JSON(fiber.Map{
			"path":    "/long",
			"message": "OK",
		})
	})

	app.Listen(":3000")
}
