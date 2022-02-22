package route

import (
	"active-order-management/internal/handler/health_check_handler"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(api fiber.Router) {

	healtCheck := api.Group("health-check")

	healtCheck.Get("ping", health_check_handler.Ping)

	healtCheck.Get("db-check",health_check_handler.DBCheck)
}
