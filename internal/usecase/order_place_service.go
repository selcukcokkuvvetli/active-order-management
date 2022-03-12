package usecase

import "active-order-management/domain/repository"

type OrderPlaceService struct {
	repository repository.Repository
}

func NewOrderPlaceService(orderPlaceRepository *repository.Repository) Service {
	return &OrderPlaceService{*orderPlaceRepository}
}

func (o *OrderPlaceService) Get(id string) (interface{}, error) {
	return o.repository.Get(id)
}

func (o *OrderPlaceService) GetAll() (interface{}, error) {
	return o.repository.GetAll()
}

func (o *OrderPlaceService) Post(NewOrderPlaceService interface{}) error {
	return o.repository.Add(NewOrderPlaceService)
}

func (o *OrderPlaceService) Put(existingOrderPlace interface{}) (interface{}, error) {
	return o.repository.Update(existingOrderPlace)
}

func (o *OrderPlaceService) Delete(id string) error {
	return o.repository.Delete(id)
}
