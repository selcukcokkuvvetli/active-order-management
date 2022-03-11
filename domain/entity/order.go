package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderStatus int

const (
	Preparing OrderStatus = iota + 1
	Serving
	Completed
)

type Order struct {
	ID           string          `json:"id"`
	OrderPlaceID string          `json:"order_place_id"`
	SubTotal     decimal.Decimal `json:"sub_total"`
	Total        decimal.Decimal `json:"total"`
	Discount     decimal.Decimal `json:"discount"`
	Status       OrderStatus     `json:"status"`
	WorkerPerson string          `json:"worker_person"`
	IsActive     bool            `json:"is_active"`
	IsDeleted    bool            `json:"is_deleted"`
	CreatedDate  time.Time       `json:"created_date"`
	ModifiedDate time.Time       `json:"modified_date"`
}
