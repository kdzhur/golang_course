package game

import (
	"fmt"
)

func StartGame() {
	game := newGame()

	game.board.PrintBoard()

	for i := 0; i < len(game.board.field)*len(game.board.field); i = i + 2 {
		for _, p := range game.players {
			playerMove := p.scanPlayerMove(game.board)
			p.modifyBoard(game.board, playerMove)
			if game.board.checkWinCondition(p.sign) {
				fmt.Printf("%v won!!!", p.name)
				game.board.PrintBoard()
				return
			}
			game.board.PrintBoard()
		}
	}
}
