package migration

import (
	"fmt"
)

func (rc *CommandContext) CreateTable(name string, columns string) error {
	query := `
				CREATE TABLE IF NOT EXISTS %s (
					%s
				);`

	query = fmt.Sprintf(query, name, columns)
	if _, err := rc.DB.Exec(query); err != nil {
		return err
	}

	return nil
}

func (rc *CommandContext) AlterTable(name, commands string) error {
	query := `
				ALTER TABLE %s
				  %s;`

	query = fmt.Sprintf(query, name, commands)
	if _, err := rc.DB.Exec(query); err != nil {
		return err
	}

	return nil
}

func (rc *CommandContext) Execute(queryCommand string) error {
	_, err := rc.DB.Exec(queryCommand)

	return err
}
