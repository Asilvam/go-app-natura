package main

import (
	user_service "github.com/Asilvam/go-app-natura.git/services/user.service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
)

func handleUser(c *fiber.Ctx) error {
	users, _ := user_service.GetUsers()
	return c.Status(fiber.StatusOK).JSON(users)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World GO!")
	})

	userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	//app.Get("/users", handleUser)

	app.Listen(":" + port)
}
