package migration

import "active-order-management/domain/entity"

func (mc *Context) order_table_create_20220305(migration entity.Migration) error {
	migrationVersion := "0.0.3"

	if migrationVersion == migration.Version && !migration.IsApplied {

		//Initial Command Table Create
		OrderTableColumns := `
					id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
					order_place_id uuid NOT NULL REFERENCES order_places(id),
					sub_total NUMERIC(5,2),
					total NUMERIC(5,2),
					discount NUMERIC(5,2),
					status INT NOT NULL CHECK (status>0 AND status<4),
					worker_person CHARACTER VARYING(255),
					is_active BIT,
					is_deleted BIT,
					created_date TIMESTAMP,
					modified_date TIMESTAMP
	`
		err := mc.CommandContext.CreateTable("orders", OrderTableColumns)

		return err
	}

	return nil
}
