package main

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Firstname string
	Lastname  string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Firstname: "jhon",
		Lastname:  "doe",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Get("/user", handleUser)

	app.Listen(":3000")
}
