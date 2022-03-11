package route

import (
	"active-order-management/domain/repository"
	"active-order-management/internal/handler/order_handler"
	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(api fiber.Router, repository *repository.Repository) {

	oHandler := order_handler.NewHandler(repository)

	order := api.Group("orders")

	order.Get("/:id", oHandler.Get)
	order.Get("/", oHandler.GetAll)
	order.Put("/", oHandler.Put)
	order.Post("/", oHandler.Post)
	order.Delete("/:id", oHandler.Delete)
}
