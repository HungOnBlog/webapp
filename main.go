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

	app.Post("/relay", func(c *fiber.Ctx) error {
		body := new(struct {
			Number int `json:"number"`
		})

		if err := c.BodyParser(body); err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"path":    "/relay",
			"message": "OK",
			"number":  body.Number,
		})
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		body := new(struct {
			Number int `json:"number"`
			Adder  int `json:"adder"`
		})

		if err := c.BodyParser(body); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"path":    "/add",
			"message": "OK",
			"result":  body.Number + body.Adder,
		})
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		body := new(struct {
			Number  int `json:"number"`
			Updater int `json:"updater"`
		})

		if err := c.BodyParser(body); err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"path":    "/update",
			"message": "OK",
			"result":  body.Number + body.Updater,
		})
	})

	app.Listen(":3000")
}
