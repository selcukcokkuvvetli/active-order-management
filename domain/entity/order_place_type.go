package entity

import "time"

type OrderPlaceType struct {
	ID           string
	Name         string
	Description  string
	IsActive     bool
	IsDeleted    bool
	CreatedDate  time.Time
	ModifiedDate time.Time
}
