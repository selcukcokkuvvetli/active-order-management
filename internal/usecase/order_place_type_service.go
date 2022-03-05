package usecase

import (
	"active-order-management/domain/repository"
)

type OrderPlaceTypeService struct {
	repository repository.Repository
}

func NewOrderPlaceTypeService(orderPlaceTypeRepository *repository.Repository) Service {
	return &OrderPlaceTypeService{*orderPlaceTypeRepository,
	}
}

func (o *OrderPlaceTypeService) Get(id string) (interface{}, error) {
	return o.repository.Get(id)
}

func (o *OrderPlaceTypeService) GetAll() (interface{}, error) {
	return o.repository.GetAll()
}

func (o *OrderPlaceTypeService) Post(newOrderPlaceType interface{}) error {
	return o.repository.Add(newOrderPlaceType)
}

func (o *OrderPlaceTypeService) Put(existingOrderPlaceType interface{}) (interface{}, error) {
	return o.repository.Update(existingOrderPlaceType)
}

func (o *OrderPlaceTypeService) Delete(id string) error {
	return o.repository.Delete(id)
}