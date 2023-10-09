// 1. Створити API для перегляду списку прав. Користувач повинен мати можливість переглядати список завдань за конкретну дату.

// Додаткові вимоги:

// • список завдань має бути збережений в оперативній пам'яті та бути доступним під час кожного запиту;

// • отримання списку завдань має здійснюватись методом GET на адресі "/tasks".

package main

import (
	"fmt"
	"http-server/cmd/task1/internal/data"
	"http-server/cmd/task1/utils/filters"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var tasks []data.Task

	app := fiber.New()

	app.Get("/tasks", func(c *fiber.Ctx) error {
		date := c.Query("date")
		if date == "" {
			return c.JSON(tasks)
		}
		filteredTasks := filters.FilterTasksByDate(date, tasks)
		return c.JSON(filteredTasks)
	})

	tasks = append(tasks, data.Task{ID: 1, Title: "Task 1", Date: "2023-10-08"})
	tasks = append(tasks, data.Task{ID: 2, Title: "Task 2", Date: "2023-10-09"})
	tasks = append(tasks, data.Task{ID: 3, Title: "Task 3", Date: "2023-10-09"})

	err := app.Listen(":3000")
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
