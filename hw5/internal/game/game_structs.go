package game

type game struct {
	players []*player
	board   *board
}

const (
	X_sign = "X"
	O_sign = "O"
)

func newGame() *game {
	return &game{
		players: []*player{
			newPlayer("Player #1", X_sign),
			newPlayer("Player #2", O_sign),
		},
		board: newBoard("_"),
	}
}
