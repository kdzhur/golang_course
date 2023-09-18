package game

import "the_game/internal/player"

type Game struct {
	name string
	// locations []*location
	Player *player.Player
}

type location struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Events      []event `json:"events"`
}

type event struct {
	Name        string   `json:"name"`
	Id          int      `json:"id"`
	Description string   `json:"description"`
	Choices     []Choise `json:"choices"`
	The_end     bool     `json:"the_end"`
}

type Choise struct {
	Description string `json:"descripion"`
	Reference   int    `json:"reference"`
}
