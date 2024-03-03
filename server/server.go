package main

import (
	"log"

	"github.com/AdityaLad2004/blog/database"
	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}
func main() {

	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Eror in sql conn")
	}

	defer sqlDB.Close()
	app := fiber.New()

	router.setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
