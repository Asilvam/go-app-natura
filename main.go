package main

import (
	user_service "github.com/Asilvam/go-app-natura.git/services/user.service"
	"github.com/gofiber/fiber/v2"
)

func handleUser(c *fiber.Ctx) error {
	users, _ := user_service.GetUsers()
	return c.Status(fiber.StatusOK).JSON(users)
}

func main() {

	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Get("/user", handleUser)

	app.Listen(":3000")
}
