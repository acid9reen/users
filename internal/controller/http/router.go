package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"users/config"
	"users/internal/controller/http/healthcheck"
)

func Setup(
	config *config.App,
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

	healthcheck.UseSubRoute(app)

	return app
}
