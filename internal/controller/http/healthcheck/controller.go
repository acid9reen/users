package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

func (router *Router) GetHealthStatus(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{})
}
