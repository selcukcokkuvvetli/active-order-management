package repository

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
	"database/sql"
	"fmt"
	"time"
)

const (
	tableNameOrderPlaceType    = "order_place_types"
	tableColumnsOrderPlaceType = "name, description, is_active, is_deleted, created_date, modified_date"
)

type OrderPlaceTypeRepository struct {
	DB *sql.DB
}

func NewOrderPlaceTypeRepository(db *sql.DB) Repository {
	return &OrderPlaceTypeRepository{DB: db}
}

func (optr *OrderPlaceTypeRepository) Get(id string) (interface{}, error) {
	orderPlaceType := new(entity.OrderPlaceType)
	query := fmt.Sprintf(GetQuery, tableNameOrderPlaceType, "id", id)

	row := optr.DB.QueryRow(query)
	row.Scan(&orderPlaceType.ID, &orderPlaceType.Name, &orderPlaceType.Description,
		&orderPlaceType.IsActive, &orderPlaceType.IsDeleted, &orderPlaceType.CreatedDate,
		&orderPlaceType.ModifiedDate)
	err := row.Err()
	return orderPlaceType, err
}

func (optr *OrderPlaceTypeRepository) GetAll() (interface{}, error) {
	orderPlaceTypes := make([]entity.OrderPlaceType, 0)

	query := fmt.Sprintf(GetAllQuery, tableNameOrderPlaceType)

	rows, err := optr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		orderPlaceType := new(entity.OrderPlaceType)
		rows.Scan(&orderPlaceType.ID, &orderPlaceType.Name, &orderPlaceType.Description,
			&orderPlaceType.IsActive, &orderPlaceType.IsDeleted, &orderPlaceType.CreatedDate,
			&orderPlaceType.ModifiedDate)
	}
	return orderPlaceTypes, nil
}

func (optr *OrderPlaceTypeRepository) Last() (interface{}, error) {
	orderPlaceType := new(entity.OrderPlaceType)
	query := fmt.Sprintf(LastQuery, tableNameOrderPlaceType, "created_date")

	row := optr.DB.QueryRow(query)
	row.Scan(&orderPlaceType.ID, &orderPlaceType.Name, &orderPlaceType.Description,
		&orderPlaceType.IsActive, &orderPlaceType.IsDeleted, &orderPlaceType.CreatedDate,
		&orderPlaceType.ModifiedDate)

	err := row.Err()
	return *orderPlaceType, err
}

func (optr *OrderPlaceTypeRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableNameOrderPlaceType, "id", id)

	_, err := optr.DB.Exec(query)

	return err
}

func (optr *OrderPlaceTypeRepository) Add(newModel interface{}) error {
	newOrderPlaceType := newModel.(entity.OrderPlaceType)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s'", newOrderPlaceType.Name, newOrderPlaceType.Description,
		global.BoolToPSQLBit(newOrderPlaceType.IsActive), global.BoolToPSQLBit(newOrderPlaceType.IsDeleted),
		newOrderPlaceType.CreatedDate.Format(time.RFC3339), newOrderPlaceType.ModifiedDate.Format(time.RFC3339))

	query := fmt.Sprintf(AddQuery, tableNameOrderPlaceType, tableColumnsOrderPlaceType, newEntityValues)
	_, err := optr.DB.Exec(query)
	return err
}

func (optr *OrderPlaceTypeRepository) Update(existingModel interface{}) (interface{}, error) {
	existingOrderPlaceType := existingModel.(entity.OrderPlaceType)
	updateEntityValues := fmt.Sprintf("name = '%s', description = '%s', is_active = '%s', is_deleted = '%s', modified_date = '%s'", existingOrderPlaceType.Name, existingOrderPlaceType.Description,
		global.BoolToPSQLBit(existingOrderPlaceType.IsActive), global.BoolToPSQLBit(existingOrderPlaceType.IsDeleted), existingOrderPlaceType.ModifiedDate.Format(time.RFC3339))
	query := fmt.Sprintf(UpdateQuery, tableNameOrderPlaceType, updateEntityValues, "id", existingOrderPlaceType.ID)

	_, err := optr.DB.Exec(query)

	return existingOrderPlaceType, err
}
