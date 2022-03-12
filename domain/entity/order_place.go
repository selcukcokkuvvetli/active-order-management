package entity

import "time"

type OrderPlace struct {
	ID           string    `json:"id"`
	TypeID       string    `json:"type_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	IsActive     bool      `json:"is_active"`
	IsDeleted    bool      `json:"is_deleted"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}
