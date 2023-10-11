// 2. Який отримує запити на адресі "/translate" методом POST та перекладає текст з однієї мови на іншу.

// Додаткові вимоги:

// • використовувати сторонній API для перекладу тексту;

// • запит до API має включати параметри, які вказують мову, з якої та на яку перекладається текст;

package main

import (
	"net/http"
	"rest/cmd/task2/internal/translate"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/translate", func(c *fiber.Ctx) error {
		type TranslateRequest struct {
			Q      string `json:"q"`
			Source string `json:"source"`
			Target string `json:"target"`
			Format string `json:"format"`
			APIKey string `json:"api_key"`
		}
		var request TranslateRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "400: Unproper input",
			})
		}

		translatedText, err := translate.TranslateText(request.Q, request.Source, request.Target, request.Format)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"translatedText": translatedText,
		})
	})

	app.Listen(":3000")
}
