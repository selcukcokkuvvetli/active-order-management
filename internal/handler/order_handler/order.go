package order_handler

import (
	"active-order-management/domain/entity"
	"active-order-management/domain/repository"
	"active-order-management/internal/usecase"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service usecase.Service
}

func NewHandler(orderRepository *repository.Repository) *Handler {
	return &Handler{service: usecase.NewOrderService(orderRepository)}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	orderId := c.Params("id")
	orderInt, err := h.service.Get(orderId)
	if err != nil {
		return err
	}
	result := orderInt.(*entity.Order)
	return c.JSON(result)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	orderInt, err := h.service.GetAll()
	if err != nil {
		return err
	}
	result := orderInt.([]entity.Order)
	return c.JSON(result)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	var newOrder entity.Order
	err := json.Unmarshal(c.Body(), &newOrder)
	if err != nil {
		return err
	}

	err = h.service.Post(newOrder)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}

func (h *Handler) Put(c *fiber.Ctx) error {
	var existingO entity.Order
	err := json.Unmarshal(c.Body(), &existingO)
	if err != nil {
		return err
	}

	serviceResult, err := h.service.Put(existingO)
	if err != nil {
		return err
	}
	return c.JSON(serviceResult)

}

func (h *Handler) Delete(c *fiber.Ctx) error {
	orderId := c.Params("id")
	err := h.service.Delete(orderId)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}
