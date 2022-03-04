package migration

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
)

func (mc *Context) Apply(migrations []entity.Migration) {
	if len(migrations) > 0 {
		for _, migration := range migrations {
			mc.order_place_type_table_create_20220223(migration)
			mc.order_place_table_create_20220303(migration)
			migration.IsApplied = true
			mc.MigrationRepository.Update(migration)
		}

	} else {
		err := mc.initial_migration_table_create_20220216(entity.Migration{Version: "0.0.0"})
		global.PanicIfError(err)
	}
}

func (mc *Context) initial_migration_table_create_20220216(migration entity.Migration) error {
	migrationVersion := "0.0.0"

	if migrationVersion == migration.Version {
		err := mc.CommandContext.Execute("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
		if err != nil {
			return err
		}
		//Initial Command Table Create
		migrationTableColumns := `
					version CHARACTER VARYING(10) PRIMARY KEY,
					description CHARACTER VARYING(255),
					is_applied bit
	`
		err = mc.CommandContext.CreateTable("migrations", migrationTableColumns)

		return err
	}

	return nil
}
