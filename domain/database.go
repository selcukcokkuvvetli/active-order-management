package domain

import (
	"active-order-management/domain/entity"
	"active-order-management/domain/migration"
	"active-order-management/domain/repository"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func (dc *DatabaseContext) Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=public", dc.Host, dc.Port, dc.UserName, dc.Password, dc.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("\n Database connection successfuly done and opened. \n")
	return db, err
}

func (dc *DatabaseContext) Close(db *sql.DB) {
	defer db.Close()
	fmt.Println("\n Database connection successfuly closed. \n")
}

func (dc *DatabaseContext) Migrate(db *sql.DB) error {

	migrationRepository := repository.NewMigrationRepository(db)
	migrationContext := migration.NewContext(db, migrationRepository)
	
	// If migrations table is not created, we must create it first
	migrationContext.Apply(nil)

	migrationRepository.Add(entity.Migration{
		Version:     "0.0.1",
		Description: "Order place type created",
		IsApplied:   false,
	})

	migrationRepository.Add(entity.Migration{
		Version:     "0.0.2",
		Description: "Order place created",
		IsApplied:   false,
	})

	migrationRepository.Add(entity.Migration{
		Version:     "0.0.3",
		Description: "Order  created",
		IsApplied:   false,
	})
  
  migrationRepository.Add(entity.Migration{
		Version:     "0.0.4",
		Description: "Order Item table created",
		IsApplied:   false,
	})

	migrationInterface, _ := migrationRepository.GetAll()
	migrations := migrationInterface.([]entity.Migration)

	migrationContext.Apply(migrations)

	return nil
}
