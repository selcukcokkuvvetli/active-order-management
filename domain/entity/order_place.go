package entity

import "time"

type OrderPlace struct {
	ID           string
	TypeID       string
	Name         string
	Description  string
	IsActive     bool
	IsDeleted    bool
	CreatedDate  time.Time
	ModifiedDate time.Time
}
