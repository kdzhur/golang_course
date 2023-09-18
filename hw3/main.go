package main

import (
	"the_game/internal/game"
	"the_game/internal/player"
)

func main() {
	p := player.NewPlayer("CHARACTER_NAME")

	game.LoadGame(p)
}
