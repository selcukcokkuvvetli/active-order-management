package repository

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
	"database/sql"
	"fmt"
	"time"
)

const (
	tableNameOrderPlace    = "order_places"
	tableColumnsOrderPlace = "name, type_id, description, is_active, is_deleted, created_date, modified_date"
)

type OrderPlaceRepository struct {
	DB *sql.DB
}

func NewOrderPlaceRepository(db *sql.DB) Repository {
	return &OrderPlaceRepository{DB: db}
}

func (opr *OrderPlaceRepository) Get(id string) (interface{}, error) {
	orderPlace := new(entity.OrderPlace)
	query := fmt.Sprintf(GetQuery, tableNameOrderPlace, "id", id)

	row := opr.DB.QueryRow(query)
	row.Scan(&orderPlace.ID, &orderPlace.TypeID, &orderPlace.Name, &orderPlace.Description,
		&orderPlace.IsActive, &orderPlace.IsDeleted, &orderPlace.CreatedDate,
		&orderPlace.ModifiedDate)
	err := row.Err()
	return orderPlace, err
}

func (opr *OrderPlaceRepository) GetAll() (interface{}, error) {
	orderPlaces := make([]entity.OrderPlace, 0)

	query := fmt.Sprintf(GetAllQuery, tableNameOrderPlace)

	rows, err := opr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		orderPlace := new(entity.OrderPlace)
		rows.Scan(&orderPlace.ID, &orderPlace.TypeID, &orderPlace.Name, &orderPlace.Description,
			&orderPlace.IsActive, &orderPlace.IsDeleted, &orderPlace.CreatedDate,
			&orderPlace.ModifiedDate)
	}
	return orderPlaces, nil
}

func (opr *OrderPlaceRepository) Last() (interface{}, error) {
	orderPlace := new(entity.OrderPlace)
	query := fmt.Sprintf(LastQuery, tableNameOrderPlace, "created_date")

	row := opr.DB.QueryRow(query)
	row.Scan(&orderPlace.ID, &orderPlace.TypeID, &orderPlace.Name, &orderPlace.Description,
		&orderPlace.IsActive, &orderPlace.IsDeleted, &orderPlace.CreatedDate,
		&orderPlace.ModifiedDate)

	err := row.Err()
	return orderPlace, err
}

func (opr *OrderPlaceRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableNameOrderPlace, "id", id)

	_, err := opr.DB.Exec(query)

	return err
}

func (opr *OrderPlaceRepository) Add(newModel interface{}) error {
	newOrderPlace := newModel.(entity.OrderPlace)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s', '%s', '%s', '%s', '%s'", newOrderPlace.Name, newOrderPlace.TypeID, newOrderPlace.Description,
		global.BoolToPSQLBit(newOrderPlace.IsActive), global.BoolToPSQLBit(newOrderPlace.IsDeleted),
		newOrderPlace.CreatedDate.Format(time.RFC3339), newOrderPlace.ModifiedDate.Format(time.RFC3339))

	query := fmt.Sprintf(AddQuery, tableNameOrderPlace, tableColumnsOrderPlace, newEntityValues)
	_, err := opr.DB.Exec(query)
	return err
}

func (opr *OrderPlaceRepository) Update(existingModel interface{}) (interface{}, error) {
	existingOrderPlace := existingModel.(entity.OrderPlace)
	updateEntityValues := fmt.Sprintf("name = '%s', type_id = '%s', description = '%s', is_active = '%s', is_deleted = '%s', modified_date = '%s'", existingOrderPlace.Name, existingOrderPlace.TypeID, existingOrderPlace.Description,
		global.BoolToPSQLBit(existingOrderPlace.IsActive), global.BoolToPSQLBit(existingOrderPlace.IsDeleted), existingOrderPlace.ModifiedDate.Format(time.RFC3339))
	query := fmt.Sprintf(UpdateQuery, tableNameOrderPlace, updateEntityValues, "id", existingOrderPlace.ID)

	_, err := opr.DB.Exec(query)

	return existingOrderPlace, err
}
