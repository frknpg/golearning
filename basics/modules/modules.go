package main

import (
	"fmt"
	. "training/sub_module"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fmt.Println("Sub module var: ", Public)

	app.Listen(":3000")
}
