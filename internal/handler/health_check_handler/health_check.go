package health_check_handler

import (
	"active-order-management/domain"
	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	return c.SendString("Pong")
}

func DBCheck(c *fiber.Ctx) error {
	dbContext := domain.NewDatabaseContext("localhost","active-order-management-db","aomdbuser","aom123.",5432)
	db,err := dbContext.Open()

	if err != nil {
		return c.SendString("Database is not running")
	} else {
		defer dbContext.Close(db)
		return c.SendString("Database is alive")
	}
}
