package migration

import (
	"active-order-management/domain/migration"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlterTable(t *testing.T) {
	const tableName = "test_table"
	const commands = "Alter Column test_id INT"
	query := `
				ALTER TABLE %s
				  %s;`
	query = fmt.Sprintf(query, tableName, commands)
	db,sqlMock,_ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	command := migration.NewCommandContext(*db)
	defer db.Close()


	t.Run("Altering table successfully", func(t *testing.T) {

		sqlMock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))

		err := command.AlterTable(tableName, commands)

		assert.Nil(t, err)
	})

	t.Run("Alter table returns error", func(t *testing.T) {
		want := errors.New("Altering table failed")
		sqlMock.ExpectExec(query).WillReturnError(want)

		got := command.AlterTable(tableName, commands)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}
