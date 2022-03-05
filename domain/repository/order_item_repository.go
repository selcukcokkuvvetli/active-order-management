package repository

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
	"database/sql"
	"fmt"
	"time"
)

const (
	tableNameOrderItem    = "order_items"
	tableColumnsOrderItem = "order_id, name, description, price, price_vat, is_active, is_deleted, created_date, modified_date"
)

type OrderItemRepository struct {
	DB *sql.DB
}

func NewOrderItemRepository(db *sql.DB) *OrderItemRepository {
	return &OrderItemRepository{DB: db}
}

func (opr *OrderItemRepository) Get(id string) (interface{}, error) {
	orderItem := new(entity.OrderItem)
	query := fmt.Sprintf(GetQuery, tableNameOrderItem, "id", id)

	row := opr.DB.QueryRow(query)
	row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
		&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
		&orderItem.CreatedDate, &orderItem.ModifiedDate)
	err := row.Err()
	return orderItem, err
}

func (opr *OrderItemRepository) GetAll() (interface{}, error) {
	orderItems := make([]entity.OrderItem, 0)

	query := fmt.Sprintf(GetAllQuery, tableNameOrderItem)

	rows, err := opr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		orderItem := new(entity.OrderItem)
		rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
			&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
			&orderItem.CreatedDate, &orderItem.ModifiedDate)

		orderItems = append(orderItems, *orderItem)
	}
	return orderItems, nil
}

func (opr *OrderItemRepository) GetAllByOrderId(orderId string) (interface{}, error) {
	orderItems := make([]entity.OrderItem, 0)

	query := fmt.Sprintf(GetQuery, tableNameOrderItem, "order_id", orderId)

	rows, err := opr.DB.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		orderItem := new(entity.OrderItem)
		rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
			&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
			&orderItem.CreatedDate, &orderItem.ModifiedDate)

		orderItems = append(orderItems, *orderItem)
	}
	return orderItems, nil
}

func (opr *OrderItemRepository) Last() (interface{}, error) {
	orderItem := new(entity.OrderItem)
	query := fmt.Sprintf(LastQuery, tableNameOrderItem, "created_date")

	row := opr.DB.QueryRow(query)
	row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
		&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
		&orderItem.CreatedDate, &orderItem.ModifiedDate)

	err := row.Err()
	return orderItem, err
}

func (opr *OrderItemRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableNameOrderItem, "id", id)

	_, err := opr.DB.Exec(query)

	return err
}

func (opr *OrderItemRepository) Add(newModel interface{}) error {
	newOrderItem := newModel.(entity.OrderItem)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s', '%s','%s', '%s', '%s', '%s', '%s'", newOrderItem.OrderID, newOrderItem.Name, newOrderItem.Description,
		newOrderItem.Price, newOrderItem.PriceVat, global.BoolToPSQLBit(newOrderItem.IsActive), global.BoolToPSQLBit(newOrderItem.IsDeleted),
		newOrderItem.CreatedDate.Format(time.RFC3339), newOrderItem.ModifiedDate.Format(time.RFC3339))

	query := fmt.Sprintf(AddQuery, tableNameOrderItem, tableColumnsOrderItem, newEntityValues)
	_, err := opr.DB.Exec(query)
	return err
}

func (opr *OrderItemRepository) Update(existingModel interface{}) (interface{}, error) {
	existingOrderItem := existingModel.(entity.OrderItem)
	updateEntityValues := fmt.Sprintf("name = '%s',  description = '%s', price = '%s', price_vat = '%s', is_active = '%s', is_deleted = '%s', modified_date = '%s'", existingOrderItem.Name, existingOrderItem.Description,
		existingOrderItem.Price, existingOrderItem.PriceVat, global.BoolToPSQLBit(existingOrderItem.IsActive),
		global.BoolToPSQLBit(existingOrderItem.IsDeleted), existingOrderItem.ModifiedDate.Format(time.RFC3339))
	query := fmt.Sprintf(UpdateQuery, tableNameOrderItem, updateEntityValues, "id", existingOrderItem.ID)

	_, err := opr.DB.Exec(query)

	return existingOrderItem, err
}
