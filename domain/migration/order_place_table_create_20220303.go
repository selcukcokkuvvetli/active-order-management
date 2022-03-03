package migration

import "active-order-management/domain/entity"

func (mc *Context) order_place_table_create_20220303(migration entity.Migration) error {
	migrationVersion := "0.0.2"

	if migrationVersion == migration.Version && !migration.IsApplied {

		//Initial Command Table Create
		OrderPlaceTableColumns := `
					id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
					CONSTRAINT type_id
      					FOREIGN KEY(id) 
	  						REFERENCES order_place_types(id),
					name CHARACTER VARYING(255),
					description CHARACTER VARYING(255),
					is_active BIT,
					is_deleted BIT,
					created_date TIMESTAMP,
					modified_date TIMESTAMP
	`
		err := mc.CommandContext.CreateTable("order_places", OrderPlaceTableColumns)

		return err
	}

	return nil
}
