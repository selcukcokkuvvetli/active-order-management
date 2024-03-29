package route

import (
	"active-order-management/domain/repository"
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")

	HealthCheckRoutes(api)

	orderPlaceTypeRepository := repository.NewOrderPlaceTypeRepository(db)
	OrderPlaceTypeRoutes(api, &orderPlaceTypeRepository)

	orderPlaceRepository := repository.NewOrderPlaceRepository(db)
	OrderPlaceRoutes(api, &orderPlaceRepository)

	orderRepository := repository.NewOrderRepository(db)
	OrderRoutes(api, &orderRepository)

	orderItemRepository := repository.NewOrderItemRepository(db)
	OrderItemRoutes(api, orderItemRepository)
}
