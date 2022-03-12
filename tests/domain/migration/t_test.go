package migration

import (
	"active-order-management/domain/migration"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {

	db,sqlMock,_ := sqlmock.New()
	command := migration.NewCommandContext(*db)
	defer db.Close()

	query := `
				ALTER TABLE %s
				  %s;`

	query = fmt.Sprintf(query, "name", "test")

	sqlMock.ExpectExec(query).WillReturnError(errors.New("Aga yanlışlık oldu"))

	err := command.AlterTable("name","test")

	assert.Nil(t, err)
}
