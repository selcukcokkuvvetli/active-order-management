package route

import (
	"active-order-management/domain/repository"
	"active-order-management/internal/handler/order_place_handler"
	"github.com/gofiber/fiber/v2"
)

func OrderPlaceRoutes(api fiber.Router, repository *repository.Repository) {

	optHandler := order_place_handler.NewHandler(repository)

	orderPlace := api.Group("order-places")

	orderPlace.Get("/:id", optHandler.Get)
	orderPlace.Get("/", optHandler.GetAll)
	orderPlace.Put("/", optHandler.Put)
	orderPlace.Post("/", optHandler.Post)
	orderPlace.Delete("/:id", optHandler.Delete)
}
