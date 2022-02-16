package migration

import (
	"active-order-management/domain/entity"
	"database/sql"
)

type Migration interface {
	Apply(migrations []entity.Migration)
}

type Context struct {
	CommandContext Command
}

func NewContext(db *sql.DB) Migration {
	return &Context{CommandContext: NewCommandContext(*db)}
}
