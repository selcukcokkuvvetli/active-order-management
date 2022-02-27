package route

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/api")

	HealthCheckRoutes(api)
}
