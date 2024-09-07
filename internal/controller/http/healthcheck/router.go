package healthcheck

import "github.com/gofiber/fiber/v2"

type Router struct {
	logger LoggerInterface
}

func UseSubRoute(group fiber.Router, logger LoggerInterface) {
	router := Router{logger}
	group.Get("/healthcheck", router.GetHealthStatus)

}
