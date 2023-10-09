package students

import (
	"fmt"
	"http-server/cmd/task2/internal/data"
)

func ParseStudentID(id string, students []data.Student) int {
	var studentID int
	_, err := fmt.Sscanf(id, "%d", &studentID)
	if err != nil {
		return -1
	}
	return studentID
}

func FindStudentByID(id int, students []data.Student) (*data.Student, error) {
	for _, student := range students {
		if student.ID == id {
			return &student, nil
		}
	}
	return nil, fmt.Errorf("student by ID: %d wasn't found", id)
}
