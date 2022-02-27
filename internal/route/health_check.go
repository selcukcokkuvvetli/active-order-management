package route

import (
	"active-order-management/internal/handler/health_check_handler"
	"github.com/gofiber/fiber/v2"
)

func HealthCheckRoutes(api fiber.Router) {

	healthCheck := api.Group("health-check")

	healthCheck.Get("ping", health_check_handler.Ping)

	healthCheck.Get("db-check",health_check_handler.DBCheck)
}
