package migration

import (
	"active-order-management/domain/entity"
	"active-order-management/domain/repository"
	"database/sql"
)

type Migration interface {
	Apply(migrations []entity.Migration)
}

type Context struct {
	CommandContext      Command
	MigrationRepository repository.Repository
}

func NewContext(db *sql.DB, mrp repository.Repository) Migration {
	return &Context{CommandContext: NewCommandContext(*db), MigrationRepository: mrp}
}
