package repository

const (
	GetQuery = `
				SELECT 
					* 
				FROM
					 %s
				WHERE
					  %s = '%s';`
	GetAllQuery = `
				SELECT 
					* 
				FROM
					 %s;`
	LastQuery = `
				SELECT
					*
				FROM
					 %s
				ORDER BY
					 %s
				DESC
				LIMIT 1;`
	DeleteQuery = `
				DELETE FROM 
							%s
				WHERE
					%s = '%s';`
	AddQuery = `
				INSERT INTO
					%s (%s)
				VALUES (
						%s
					   );`
	UpdateQuery = `
				UPDATE 
					%s
				SET 
					%s
				WHERE
					%s = '%s';
`
)

type Repository interface {
	Get(id string) (interface{}, error)
	GetAll() (interface{}, error)
	Last() (interface{}, error)
	Delete(id string) error
	Add(interface{}) error
	Update(interface{}) (interface{}, error)
}