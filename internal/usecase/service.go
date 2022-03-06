package usecase

type Service interface {
	Get(id string) (interface{}, error)
	GetAll() (interface{}, error)
	Post(interface{}) error
	Put(interface{}) (interface{}, error)
	Delete(id string) error
}