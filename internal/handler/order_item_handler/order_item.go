package order_item

import (
	"active-order-management/domain/entity"
	"active-order-management/domain/repository"
	"active-order-management/internal/usecase"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service usecase.OrderItemService
}

func NewHandler(orderItemRepository *repository.OrderItemRepository) *Handler {
	return &Handler{service: usecase.NewOrderItemService(orderItemRepository)}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	orderItemId := c.Params("id")
	orderItemInt, err := h.service.Get(orderItemId)
	if err != nil {
		return err
	}
	result := orderItemInt.(*entity.OrderItem)
	return c.JSON(result)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	orderItemsInt, err := h.service.GetAll()
	if err != nil {
		return err
	}
	result := orderItemsInt.([]entity.OrderItem)
	return c.JSON(result)
}

func (h *Handler) GetAllByOrderId(c *fiber.Ctx) error {
	orderID := c.Params("order_id")
	orderItemsInt, err := h.service.GetAllByOrderId(orderID)
	if err != nil {
		return err
	}
	result := orderItemsInt.([]entity.OrderItem)
	return c.JSON(result)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	var newOrderItem entity.OrderItem
	err := json.Unmarshal(c.Body(), &newOrderItem)
	if err != nil {
		return err
	}

	err = h.service.Post(newOrderItem)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}

func (h *Handler) Put(c *fiber.Ctx) error {
	var existingOrderItem entity.OrderItem
	err := json.Unmarshal(c.Body(), &existingOrderItem)
	if err != nil {
		return err
	}

	serviceResult, err := h.service.Put(existingOrderItem)
	if err != nil {
		return err
	}
	return c.JSON(serviceResult)

}

func (h *Handler) Delete(c *fiber.Ctx) error {
	orderItemId := c.Params("id")
	err := h.service.Delete(orderItemId)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}
