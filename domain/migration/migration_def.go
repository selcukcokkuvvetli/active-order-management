package migration

import "database/sql"

type Migration interface {
	CreateTable(name, columns string) error
	AlterTable(name, commands string) error
}

type Context struct {
	DB sql.DB
}

func NewContext(db sql.DB) Migration {
	return &Context{DB: db}
}
