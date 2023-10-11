// 1. Який отримує запити на адресі "/weather" методом GET та повертає погодні дані для заданого міста.

// Додаткові вимоги:

//     ◦ використовувати сторонній API, який надає погодні дані;

//     ◦ запит до API має включати параметр, який вказує місто, для якого повертаються погодні дані;

//     ◦ результат повертати у вигляді JSON об'єкту, де ключі — це погодні параметри (температура, вітер, вологість тощо), а значення — це відповідні значення.

package main

import (
	"net/http"
	"rest/cmd/task1/internal/weather"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/weather", func(c *fiber.Ctx) error {
		city := c.Query("city")

		if city == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Параметр city відсутній або некоректний",
			})
		}

		weatherData, err := weather.GetWeather(city)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(weatherData)
	})

	app.Listen(":3000")
}
