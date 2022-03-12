package route

import (
	"active-order-management/domain/repository"
	"active-order-management/internal/handler/order_place_handler"
	"github.com/gofiber/fiber/v2"
)

func OrderPlaceRoutes(api fiber.Router, repository *repository.Repository) {

	opHandler := order_place_handler.NewHandler(repository)

	orderPlace := api.Group("order-places")

	orderPlace.Get("/:id", opHandler.Get)
	orderPlace.Get("/", opHandler.GetAll)
	orderPlace.Put("/", opHandler.Put)
	orderPlace.Post("/", opHandler.Post)
	orderPlace.Delete("/:id", opHandler.Delete)
}
