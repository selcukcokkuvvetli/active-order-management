package route

import (
	"github.com/gofiber/fiber/v2"
)

func InitilizeRoutes(app *fiber.App) {
	api := app.Group("/api")

	HealthCheckRoutes(api)
}
