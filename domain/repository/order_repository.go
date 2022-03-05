package repository

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
	"database/sql"
	"fmt"
	"time"
)

const (
	tableNameOrder    = "orders"
	tableColumnsOrder = "order_place_id, sub_total, total, discount, status, worker_person, is_active, is_deleted, created_date, modified_date"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) Repository {
	return &OrderRepository{DB: db}
}

func (or *OrderRepository) Get(id string) (interface{}, error) {
	order := new(entity.Order)
	query := fmt.Sprintf(GetQuery, tableNameOrder, "id", id)

	row := or.DB.QueryRow(query)
	row.Scan(&order.ID, &order.OrderPlaceID, &order.SubTotal, &order.Total,
		&order.Discount, &order.Status, &order.WorkerPerson,
		&order.IsActive, &order.IsDeleted, &order.CreatedDate, &order.ModifiedDate)
	err := row.Err()
	return order, err
}

func (or *OrderRepository) GetAll() (interface{}, error) {
	orders := make([]entity.Order, 0)

	query := fmt.Sprintf(GetAllQuery, tableNameOrder)

	rows, err := or.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		order := new(entity.Order)
		rows.Scan(&order.ID, &order.OrderPlaceID, &order.SubTotal, &order.Total,
			&order.Discount, &order.Status, &order.WorkerPerson,
			&order.IsActive, &order.IsDeleted, &order.CreatedDate, &order.ModifiedDate)
		orders = append(orders, *order)
	}
	return orders, nil
}

func (or *OrderRepository) Last() (interface{}, error) {
	order := new(entity.Order)
	query := fmt.Sprintf(LastQuery, tableNameOrder, "created_date")

	row := or.DB.QueryRow(query)
	row.Scan(&order.ID, &order.OrderPlaceID, &order.SubTotal, &order.Total,
		&order.Discount, &order.Status, &order.WorkerPerson,
		&order.IsActive, &order.IsDeleted, &order.CreatedDate, &order.ModifiedDate)

	err := row.Err()
	return order, err
}

func (or *OrderRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableNameOrder, "id", id)

	_, err := or.DB.Exec(query)

	return err
}

func (or *OrderRepository) Add(newModel interface{}) error {
	newOrder := newModel.(entity.Order)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s', '%s', '%d', '%s', '%s', '%s', '%s', '%s'", newOrder.OrderPlaceID, newOrder.SubTotal, newOrder.Total,
		newOrder.Discount, newOrder.Status, newOrder.WorkerPerson, global.BoolToPSQLBit(newOrder.IsActive), global.BoolToPSQLBit(newOrder.IsDeleted),
		newOrder.CreatedDate.Format(time.RFC3339), newOrder.ModifiedDate.Format(time.RFC3339))

	query := fmt.Sprintf(AddQuery, tableNameOrder, tableColumnsOrder, newEntityValues)
	_, err := or.DB.Exec(query)
	return err
}

func (or *OrderRepository) Update(existingModel interface{}) (interface{}, error) {
	existingOrder := existingModel.(entity.Order)
	updateEntityValues := fmt.Sprintf("sub_total = '%s', total = '%s', discount = '%s', status = '%d', worker_person = '%s',  is_active = '%s', is_deleted = '%s', modified_date = '%s'",
		existingOrder.SubTotal, existingOrder.Total, existingOrder.Discount, existingOrder.Status, existingOrder.WorkerPerson,
		global.BoolToPSQLBit(existingOrder.IsActive), global.BoolToPSQLBit(existingOrder.IsDeleted), existingOrder.ModifiedDate.Format(time.RFC3339))
	query := fmt.Sprintf(UpdateQuery, tableNameOrder, updateEntityValues, "id", existingOrder.ID)

	_, err := or.DB.Exec(query)

	return existingOrder, err
}
