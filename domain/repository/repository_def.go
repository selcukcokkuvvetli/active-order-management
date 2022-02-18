package repository

const (
	GetQuery = `
				SELECT 
					* 
				FROM
					 %s
				WHERE
					  %s = %s;`
	GetAllQuery = `
				SELECT 
					* 
				FROM
					 %s;`
)

type Repository interface {
	Get(id string) (interface{}, error)
	GetAll() ([]interface{}, error)
	Last() (interface{}, error)
	Delete(id string) error
	Add(interface{}) error
	Update(interface{}) (interface{}, error)
}