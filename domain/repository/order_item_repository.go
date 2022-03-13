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

func (oir *OrderItemRepository) Get(id string) (interface{}, error) {
	orderItem := new(entity.OrderItem)
	query := fmt.Sprintf(GetQuery, tableNameOrderItem, "id", id)

	row := oir.DB.QueryRow(query)
	row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
		&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
		&orderItem.CreatedDate, &orderItem.ModifiedDate)
	err := row.Err()
	return orderItem, err
}

func (oir *OrderItemRepository) GetAll() (interface{}, error) {
	orderItems := make([]entity.OrderItem, 0)

	query := fmt.Sprintf(GetAllQuery, tableNameOrderItem)

	rows, err := oir.DB.Query(query)
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

func (oir *OrderItemRepository) GetAllByOrderId(orderId string) (interface{}, error) {
	orderItems := make([]entity.OrderItem, 0)

	query := fmt.Sprintf(GetQuery, tableNameOrderItem, "order_id", orderId)
	fmt.Println(orderId)
	rows, err := oir.DB.Query(query)

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
	return orderItems, err
}

func (oir *OrderItemRepository) Last() (interface{}, error) {
	orderItem := new(entity.OrderItem)
	query := fmt.Sprintf(LastQuery, tableNameOrderItem, "created_date")

	row := oir.DB.QueryRow(query)
	row.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.Name, &orderItem.Description,
		&orderItem.Price, &orderItem.PriceVat, &orderItem.IsActive, &orderItem.IsDeleted,
		&orderItem.CreatedDate, &orderItem.ModifiedDate)

	err := row.Err()
	return orderItem, err
}

func (oir *OrderItemRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableNameOrderItem, "id", id)

	_, err := oir.DB.Exec(query)

	return err
}

func (oir *OrderItemRepository) Add(newModel interface{}) error {
	newOrderItem := newModel.(entity.OrderItem)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s', '%s','%s', '%s', '%s', '%s', '%s'", newOrderItem.OrderID, newOrderItem.Name, newOrderItem.Description,
		newOrderItem.Price, newOrderItem.PriceVat, global.BoolToPSQLBit(newOrderItem.IsActive), global.BoolToPSQLBit(newOrderItem.IsDeleted),
		newOrderItem.CreatedDate.Format(time.RFC3339), newOrderItem.ModifiedDate.Format(time.RFC3339))

	query := fmt.Sprintf(AddQuery, tableNameOrderItem, tableColumnsOrderItem, newEntityValues)
	_, err := oir.DB.Exec(query)
	return err
}

func (oir *OrderItemRepository) Update(existingModel interface{}) (interface{}, error) {
	existingOrderItem := existingModel.(entity.OrderItem)
	updateEntityValues := fmt.Sprintf("name = '%s',  description = '%s', price = '%s', price_vat = '%s', is_active = '%s', is_deleted = '%s', modified_date = '%s'", existingOrderItem.Name, existingOrderItem.Description,
		existingOrderItem.Price, existingOrderItem.PriceVat, global.BoolToPSQLBit(existingOrderItem.IsActive),
		global.BoolToPSQLBit(existingOrderItem.IsDeleted), existingOrderItem.ModifiedDate.Format(time.RFC3339))
	query := fmt.Sprintf(UpdateQuery, tableNameOrderItem, updateEntityValues, "id", existingOrderItem.ID)

	_, err := oir.DB.Exec(query)

	return existingOrderItem, err
}
