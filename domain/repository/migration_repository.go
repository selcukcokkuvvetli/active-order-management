package repository

import (
	"active-order-management/domain/entity"
	"active-order-management/global"
	"database/sql"
	"fmt"
)

const (
	tableName    = "migrations"
	tableColumns = "version, description, is_applied"
)

type MigrationRepository struct {
	DB *sql.DB
}

func NewMigrationRepository(db *sql.DB) Repository {
	return &MigrationRepository{DB: db}
}

// Get will return a migration with given id. Eg: entity.migration
func (mr *MigrationRepository) Get(id string) (interface{}, error) {
	migration := new(entity.Migration)
	query := fmt.Sprintf(GetQuery, tableName, "version", id)

	row := mr.DB.QueryRow(query)
	row.Scan(&migration.Version, &migration.Description, &migration.IsApplied)

	err := row.Err()
	return migration, err
}

// GetAll will always be returning slice of entity.migration. Eg: []entity.migration
func (mr *MigrationRepository) GetAll() (interface{}, error) {
	migrations := make([]entity.Migration, 0)

	query := fmt.Sprintf(GetAllQuery, tableName)

	rows, err := mr.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		migration := new(entity.Migration)
		rows.Scan(&migration.Version, &migration.Description, &migration.IsApplied)

		migrations = append(migrations, *migration)
	}

	return migrations, nil
}

// Last will return the last migration record in migration table
func (mr *MigrationRepository) Last() (interface{}, error) {
	migration := new(entity.Migration)
	query := fmt.Sprintf(LastQuery, tableName, "version")

	row := mr.DB.QueryRow(query)
	row.Scan(&migration.Version, &migration.Description, &migration.IsApplied)

	err := row.Err()
	return *migration, err
}

// Delete will delete a migration record with given id
func (mr *MigrationRepository) Delete(id string) error {
	query := fmt.Sprintf(DeleteQuery, tableName, "version", id)

	_, err := mr.DB.Exec(query)

	return err
}

// Add will insert new migration record
func (mr *MigrationRepository) Add(newModel interface{}) error {
	newMigration := newModel.(entity.Migration)
	newEntityValues := fmt.Sprintf("'%s', '%s', '%s'", newMigration.Version, newMigration.Description, global.BoolToPSQLBit(newMigration.IsApplied))
	query := fmt.Sprintf(AddQuery, tableName, tableColumns, newEntityValues)

	_, err := mr.DB.Exec(query)

	return err
}

// Update will update a migration with given migration record and return updated migration
func (mr *MigrationRepository) Update(existingModel interface{}) (interface{}, error) {
	existingMigration := existingModel.(entity.Migration)
	updateEntityValues := fmt.Sprintf("description = '%s', is_applied = '%s'", existingMigration.Description, global.BoolToPSQLBit(existingMigration.IsApplied))
	query := fmt.Sprintf(UpdateQuery, tableName, updateEntityValues, "version", existingMigration.Version)

	_, err := mr.DB.Exec(query)

	return existingMigration, err
}
