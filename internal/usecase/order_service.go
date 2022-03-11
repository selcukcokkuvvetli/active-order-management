package usecase

import "active-order-management/domain/repository"

type OrderService struct {
	repository repository.Repository
}

func NewOrderService(orderRepository *repository.Repository) Service {
	return &OrderService{*orderRepository}
}

func (o *OrderService) Get(id string) (interface{}, error) {
	return o.repository.Get(id)
}

func (o *OrderService) GetAll() (interface{}, error) {
	return o.repository.GetAll()
}

func (o *OrderService) Post(newOrder interface{}) error {
	return o.repository.Add(newOrder)
}

func (o *OrderService) Put(existingOrder interface{}) (interface{}, error) {
	return o.repository.Update(existingOrder)
}

func (o *OrderService) Delete(id string) error {
	return o.repository.Delete(id)
}
