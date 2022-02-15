package migration

import "fmt"

func (rc *Context) CreateTable(name string, columns string) error {
	query := `
				CREATE TABLE IF NOT EXISTS %s (
					%s
				);`
	query = fmt.Sprintf(query, name, columns)
	if _, err:= rc.DB.Exec(query); err != nil {
		return err
	}

	return nil
}

func (rc *Context) AlterTable(name, commands string) error {
	 query := `
				ALTER TABLE %s
				  %s;`

	if _, err := rc.DB.Exec(query); err != nil {
		return err
	}

	return nil
}
