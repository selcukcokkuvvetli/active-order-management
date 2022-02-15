package main

import (
	"active-order-management/domain"
	"fmt"
)

func main() {
	fmt.Println("project start")
	dbContext := domain.NewDatabaseContext("localhost","active-order-management-db","aomdbuser","aom123.",5432)
	db,_ := dbContext.Open()
	dbContext.Migrate(db)
	dbContext.Close(db)
}
