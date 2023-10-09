package data

type Student struct {
	ID     int            `json:"id"`
	Name   string         `json:"name"`
	Grades map[string]int `json:"grades"`
}
