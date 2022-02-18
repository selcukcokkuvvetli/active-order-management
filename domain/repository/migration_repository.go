package repository

import (
	"active-order-management/domain/entity"
	"database/sql"
	"fmt"
)

type MigrationRepository struct {
	DB *sql.DB
}

func NewMigrationRepository(db *sql.DB) Repository {
	return &MigrationRepository{DB: db}
}

func (mr *MigrationRepository) Get(id string) (interface{}, error) {
	migration := new(entity.Migration)
	query := fmt.Sprintf(GetQuery, "migrations", "version", id)

	row := mr.DB.QueryRow(query)
	row.Scan(&migration.Version, &migration.Description, migration.IsApplied)

	err := row.Err()
	return migration, err
}

func (mr *MigrationRepository) GetAll() ([]interface{}, error) {
	migrations := make([]entity.Migration, 0)

	query := fmt.Sprintf(GetAllQuery, "migrations")

	rows, err := mr.DB.Query(query)

	for rows.Next() {
		migration := new(entity.Migration)
		rows.Scan(&migration.Version, &migration.Description, migration.IsApplied)

		migrations = append(migrations, *migration)
	}

	return mr.migrationSliceToInterfaceSlice(migrations), err
}

func (mr *MigrationRepository) Last() (interface{}, error)              { return nil, nil }
func (mr *MigrationRepository) Delete(id string) error                  { return nil }
func (mr *MigrationRepository) Add(interface{}) error                   { return nil }
func (mr *MigrationRepository) Update(interface{}) (interface{}, error) { return nil, nil }

func (mr *MigrationRepository) migrationSliceToInterfaceSlice(migrations []entity.Migration) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(migrations))
	for i, d := range migrations {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
