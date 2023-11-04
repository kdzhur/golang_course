package main

import (
	"log"
	"mvc/internal/initialize"
	"mvc/internal/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fruits := new(models.Fruits)
	app := fiber.New()

	initialize.InitApp(fruits, app)
	initialize.InitRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(fiber.ErrInternalServerError, ":", err)
	}
}
