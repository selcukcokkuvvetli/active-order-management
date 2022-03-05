package order_place_type_handler

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

func NewHandler(orderPlaceTypeRepository *repository.Repository) *Handler  {
	return &Handler{service: usecase.NewOrderPlaceTypeService(orderPlaceTypeRepository)}
}

func (h *Handler) Get(c *fiber.Ctx) error {
	orderPlaceTypeId := c.Params("id")
	orderPlaceTypeInt, err := h.service.Get(orderPlaceTypeId)
	if err != nil {
		return err
	}
	result := orderPlaceTypeInt.(*entity.OrderPlaceType)
	return c.JSON(result)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	orderPlaceTypeInt, err := h.service.GetAll()
	if err != nil {
		return err
	}
	result := orderPlaceTypeInt.([]entity.OrderPlaceType)
	return c.JSON(result)
}

func (h *Handler) Post(c *fiber.Ctx) error {
	var newOrderPlaceType entity.OrderPlaceType
	err := json.Unmarshal(c.Body(), &newOrderPlaceType)
	if err != nil {
		return err
	}

	err = h.service.Post(newOrderPlaceType)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}

func (h *Handler) Put(c *fiber.Ctx) error {
	var existingOPT entity.OrderPlaceType
	err := json.Unmarshal(c.Body(), &existingOPT)
	if err != nil {
		return err
	}

	serviceResult, err := h.service.Put(existingOPT)
	if err != nil {
		return err
	}
	return c.JSON(serviceResult)

}

func (h *Handler) Delete(c *fiber.Ctx) error {
	orderPlaceTypeId := c.Params("id")
	err := h.service.Delete(orderPlaceTypeId)
	if err != nil {
		return err
	}

	return c.SendStatus(204)
}
