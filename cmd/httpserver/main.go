package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

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

	app.Use(fiberLogger.New(fiberLogger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))

	app.Use(favicon.New(favicon.Config{
		File: "./static/img/favicon.ico",
		URL:  "/favicon.ico",
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{})
	})

	app.Listen(":" + config.HTTP.Port)
}
