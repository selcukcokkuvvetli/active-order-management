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
	ID           string
	OrderPlaceID string
	SubTotal     decimal.Decimal
	Total        decimal.Decimal
	Discount     decimal.Decimal
	Status       OrderStatus
	WorkerPerson string
	IsActive     bool
	IsDeleted    bool
	CreatedDate  time.Time
	ModifiedDate time.Time
}
