package route

import (
	"active-order-management/domain/repository"
	"active-order-management/internal/handler/order_item_handler"
	"github.com/gofiber/fiber/v2"
)

func OrderItemRoutes(api fiber.Router, repository *repository.OrderItemRepository) {

	oiHandler := order_item.NewHandler(repository)

	order := api.Group("order-items")

	order.Get("/:id", oiHandler.Get)
	order.Get("/", oiHandler.GetAll)
	api.Get("/:order_id/order-items", oiHandler.GetAllByOrderId)
	order.Put("/", oiHandler.Put)
	order.Post("/", oiHandler.Post)
	order.Delete("/:id", oiHandler.Delete)
}
