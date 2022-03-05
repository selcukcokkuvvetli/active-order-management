package migration

import "active-order-management/domain/entity"

func (mc *Context) order_item_table_creation_20220305(migration entity.Migration) error {
	migrationVersion := "0.0.4"

	if migrationVersion == migration.Version && !migration.IsApplied {

		//Initial Command Table Create
		orderItemColumns := `
					id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
					order_id uuid NOT NULL REFERENCES orders(id),
					name CHARACTER VARYING(255),
					description CHARACTER VARYING(255),
					price NUMERIC(5, 2),
					price_vat NUMERIC(5, 2),
					is_active BIT,
					is_deleted BIT,
					created_date TIMESTAMP,
					modified_date TIMESTAMP
	`
		err := mc.CommandContext.CreateTable("order_items", orderItemColumns)

		return err
	}

	return nil
}
