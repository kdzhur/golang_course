package filters

import "http-server/cmd/task1/internal/data"

func FilterTasksByDate(date string, tasks []data.Task) []data.Task {
	filtered := make([]data.Task, 0)
	for _, task := range tasks {
		if task.Date == date {
			filtered = append(filtered, task)
		}
	}
	return filtered
}
