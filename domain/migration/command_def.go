package migration

import (
	"database/sql"
)

type Command interface {
	CreateTable(name, columns string) error
	AlterTable(name, commands string) error
}

type CommandContext struct {
	DB sql.DB
}

func NewCommandContext(db sql.DB) Command {
	return &CommandContext{DB: db}
}