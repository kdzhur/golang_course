package initialize

import (
	"mvc/internal/controllers"
	"mvc/internal/models"
	"mvc/utils/generator"

	"github.com/gofiber/fiber/v2"
)

func InitApp(fruits *models.Fruits, app *fiber.App) {
	fruitList := []string{"Apple", "Banana", "Orange", "Strawberry", "Mango", "Grapes", "Cherry", "Pineapple", "Watermelon", "Kiwi"}
	fruits.Fruits = generator.GenerateRandomFruits(fruitList)

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("fruits", fruits)
		return c.Next()
	})
}

func InitRoutes(app *fiber.App) {
	app.Get("/", controllers.ListAllFruits, controllers.GetFruit)
	app.Get("/fruit", controllers.GetFruit)
}
