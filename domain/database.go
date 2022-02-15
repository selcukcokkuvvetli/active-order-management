package domain

import (
	"active-order-management/domain/migration"
	"active-order-management/global"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func (dc *DatabaseContext) Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=public", dc.Host, dc.Port, dc.UserName, dc.Password, dc.Name)
	db, err := sql.Open("postgres", psqlInfo)
	global.PanicIfError(err)

	err = db.Ping()
	global.PanicIfError(err)

	fmt.Println("\n Database connection successfuly done and opened. \n")
	return db, err
}

func (dc *DatabaseContext) Close(db *sql.DB) {
	defer db.Close()
	fmt.Println("\n Database connection successfuly closed. \n")
}

func  (dc *DatabaseContext) Migrate(db *sql.DB) error {
	migrationContext := migration.NewContext(*db)

	//Initial Migration Table Create
	migrationTableColumns := `
					version CHARACTER VARYING(10) PRIMARY KEY,
					description CHARACTER VARYING(255)
	`
	migrationContext.CreateTable("migration", migrationTableColumns)

	return nil
}