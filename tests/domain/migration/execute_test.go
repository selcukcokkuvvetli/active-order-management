package migration

import (
	"active-order-management/domain/migration"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExecute(t *testing.T) {
	commandString := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	db,sqlMock,_ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	command := migration.NewCommandContext(*db)
	defer db.Close()


	t.Run("Executes command successfully", func(t *testing.T) {

		sqlMock.ExpectExec(commandString).WillReturnResult(sqlmock.NewResult(1, 1))

		err := command.Execute(commandString)

		assert.Nil(t, err)
	})

	t.Run("Execute returns error", func(t *testing.T) {
		want := errors.New("Execute command failed")
		sqlMock.ExpectExec(commandString).WillReturnError(want)

		got := command.Execute(commandString)

		assert.NotNil(t, got)
		assert.Equal(t, want, got)
	})
}
