package route

import (
	"active-order-management/domain/repository"
	"active-order-management/internal/handler/order_place_type_handler"
	"github.com/gofiber/fiber/v2"
)

func OrderPlaceTypeRoutes(api fiber.Router, repository *repository.Repository) {

	optHandler := order_place_type_handler.NewHandler(repository)

	orderPlaceType := api.Group("order-place-types")

	orderPlaceType.Get("/:id", optHandler.Get)
	orderPlaceType.Get("/", optHandler.GetAll)
	orderPlaceType.Put("/", optHandler.Put)
	orderPlaceType.Post("/", optHandler.Post)
	orderPlaceType.Delete("/:id", optHandler.Delete)
}

