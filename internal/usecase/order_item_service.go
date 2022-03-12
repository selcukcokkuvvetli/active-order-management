package usecase

import "active-order-management/domain/repository"

type OrderItemService struct {
	repository repository.OrderItemRepository
}

func NewOrderItemService(orderItemRepository *repository.OrderItemRepository) OrderItemService {
	return OrderItemService{*orderItemRepository}
}

func (oi *OrderItemService) Get(id string) (interface{}, error) {
	return oi.repository.Get(id)
}

func (oi *OrderItemService) GetAll() (interface{}, error) {
	return oi.repository.GetAll()
}
func (oi *OrderItemService) GetAllByOrderId(orderId string) (interface{}, error) {
	return oi.repository.GetAllByOrderId(orderId)
}

func (oi *OrderItemService) Post(newOrderItem interface{}) error {
	return oi.repository.Add(newOrderItem)
}

func (oi *OrderItemService) Put(existingOrderItem interface{}) (interface{}, error) {
	return oi.repository.Update(existingOrderItem)
}

func (oi *OrderItemService) Delete(id string) error {
	return oi.repository.Delete(id)
}
