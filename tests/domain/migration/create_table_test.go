package migration

import (
	"active-order-management/domain/migration"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTable(t *testing.T) {
	const tableName = "test_table"
	const tableColumns = "test_id int primary key not null"
	query := `
				CREATE TABLE IF NOT EXISTS %s (
					%s
				);`
	query = fmt.Sprintf(query, tableName, tableColumns)
	db,sqlMock,_ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	command := migration.NewCommandContext(*db)
	defer db.Close()


	t.Run("Altering table successfully", func(t *testing.T) {

		sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

		err := command.CreateTable(tableName, tableColumns)

		assert.Nil(t, err)
	})

	t.Run("Alter table returns error", func(t *testing.T) {
		want := errors.New("Altering table failed")
		sqlMock.ExpectExec(query).WillReturnError(want)

		got := command.CreateTable(tableName, tableColumns)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}