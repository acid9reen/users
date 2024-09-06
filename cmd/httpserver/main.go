package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"

	"users/config"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Failed to load config: %e", err)
		return
	}

	app := fiber.New(fiber.Config{
		AppName: config.Name + " " + config.Version,
	})

	app.Use(favicon.New(favicon.Config{
		File: "./static/img/favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	app.Listen(":" + config.HTTP.Port)
}
