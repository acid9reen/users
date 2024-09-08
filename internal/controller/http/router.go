package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/acid9reen/users/config"
	"github.com/acid9reen/users/internal/controller/http/healthcheck"
)

func Setup(
	config *config.App,
	logger LoggerInterface,
) *fiber.App {
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

	healthcheck.UseSubRoute(app, logger)

	return app
}
