package order_place_handler

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

func NewHandler(orderPlaceRepository *repository.Repository) *Handler {
	return &Handler{service: usecase.NewOrderPlaceService(orderPlaceRepository)}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	orderPlaceId := c.Params("id")
	orderPlaceInt, err := h.service.Get(orderPlaceId)
	if err != nil {
		return err
	}
	result := orderPlaceInt.(*entity.OrderPlace)
	return c.JSON(result)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	orderPlaceInt, err := h.service.GetAll()
	if err != nil {
		return err
	}
	result := orderPlaceInt.([]entity.OrderPlace)
	return c.JSON(result)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	var newOrderPlace entity.OrderPlace
	err := json.Unmarshal(c.Body(), &newOrderPlace)
	if err != nil {
		return err
	}

	err = h.service.Post(newOrderPlace)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}

func (h *Handler) Put(c *fiber.Ctx) error {

	var existingOP entity.OrderPlace
	err := json.Unmarshal(c.Body(), &existingOP)
	if err != nil {
		return err
	}

	serviceResult, err := h.service.Put(existingOP)
	if err != nil {
		return err
	}
	return c.JSON(serviceResult)

}

func (h *Handler) Delete(c *fiber.Ctx) error {
	orderPlaceId := c.Params("id")
	err := h.service.Delete(orderPlaceId)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}
