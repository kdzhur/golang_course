// 2. Створити вебсервер для перегляду інформації щодо класу школи. Користувач повинен мати можливість отримувати загальну статистику про клас.

// Додаткові вимоги:

// • інформація про учнів має зберігатися в оперативній пам'яті та бути доступною під час кожного запиту;

// • отримання інформації про учня має здійснюватись методом GET на адресі "/student/{id}", де {id} — унікальний ідентифікатор учня;

// • дані можна отримати, лише якщо користувач є вчителем у цьому класі.
package main

import (
	"http-server/cmd/task2/internal/data"
	"http-server/cmd/task2/internal/middleware"
	"http-server/cmd/task2/utils/students"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var stdts []data.Student

	app := fiber.New()

	stdts = append(stdts, data.Student{ID: 1, Name: "Student 1", Grades: map[string]int{
		"math":               5,
		"chemistry":          4,
		"english":            5,
		"compouter sciences": 5,
	}})
	stdts = append(stdts, data.Student{ID: 2, Name: "Student 2", Grades: map[string]int{
		"math":      3,
		"chemistry": 3,
		"ukrainian": 5,
	}})
	stdts = append(stdts, data.Student{ID: 3, Name: "Student 3", Grades: map[string]int{
		"math":       5,
		"chemistry":  4,
		"literature": 3,
	}})

	app.Get("/class/statistics", func(c *fiber.Ctx) error {
		totalStudents := len(stdts)

		statistics := map[string]interface{}{
			"Total students": totalStudents,
			"Students":       stdts,
		}

		return c.JSON(statistics)
	})

	app.Get("/student/:id", middleware.Authorization(), func(c *fiber.Ctx) error {
		studentID := c.Params("id")
		id := students.ParseStudentID(studentID, stdts)
		if id == -1 {
			return fiber.ErrBadRequest
		}

		student, err := students.FindStudentByID(id, stdts)
		if err != nil {
			return fiber.ErrNotFound
		}

		return c.JSON(student)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
