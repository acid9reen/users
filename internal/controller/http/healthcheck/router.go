package healthcheck

import "github.com/gofiber/fiber/v2"

type Router struct {
}

func UseSubRoute(group fiber.Router) {
	router := Router{}
	group.Get("/healthcheck", router.GetHealthStatus)

}
