package migration

import (
	"active-order-management/domain/entity"
	"database/sql"
)

type Migration interface {
	Apply(migrations []entity.Migration)
}

type Context struct {
	DB *sql.DB
}

func NewContext(db *sql.DB) Migration {
	return &Context{DB: db}
}
