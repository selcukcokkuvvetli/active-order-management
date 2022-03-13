package entity

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderItem struct {
	ID           string          `json:"id"`
	OrderID      string          `json:"order_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Price        decimal.Decimal `json:"price"`
	PriceVat     decimal.Decimal `json:"price_vat"`
	IsActive     bool            `json:"is_active"`
	IsDeleted    bool            `json:"is_deleted"`
	CreatedDate  time.Time       `json:"created_date"`
	ModifiedDate time.Time       `json:"modified_date"`
}
