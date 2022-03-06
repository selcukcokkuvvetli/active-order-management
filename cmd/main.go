package main

import (
	"active-order-management/domain"
	"active-order-management/internal/route"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("project start")
	app := fiber.New()
	dbContext := domain.NewDatabaseContext("localhost","active-order-management-db","aomdbuser","aom123.",5432)
	db,err := dbContext.Open()
	if err == nil {
		dbContext.Migrate(db)
		defer dbContext.Close(db)
	}

	route.InitializeRoutes(app,db)

	app.Listen(":3000")
}
