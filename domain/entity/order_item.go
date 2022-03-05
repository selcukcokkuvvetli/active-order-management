package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderItem struct {
	ID           string
	OrderID      string
	Name         string
	Description  string
	Price        decimal.Decimal
	PriceVat     decimal.Decimal
	IsActive     bool
	IsDeleted    bool
	CreatedDate  time.Time
	ModifiedDate time.Time
}
