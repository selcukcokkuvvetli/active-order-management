package migration

import "active-order-management/domain/entity"

func (mc *Context) order_place_type_table_create_20220223(migration entity.Migration) error {
	migrationVersion := "0.0.1"

	if migrationVersion == migration.Version && !migration.IsApplied {

		//Initial Command Table Create
		OrderPlaceTypeTableColumns := `
					id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
					name CHARACTER VARYING(255),
					description CHARACTER VARYING(255),
					is_active BIT,
					is_deleted BIT,
					created_date TIMESTAMP,
					modified_date TIMESTAMP
	`
		err := mc.CommandContext.CreateTable("order_place_types", OrderPlaceTypeTableColumns)

		return err
	}

	return nil
}
