package game

import (
	"fmt"
	"strings"
)

// Pretty print tic-tac-toe board
func (b *board) PrintBoard() {
	fmt.Println()
	for i, row := range b.field {
		for j, column := range row {
			fmt.Print(" ", column, " ")
			if j+1 < len(b.field) {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i+1 < len(b.field) {
			fmt.Println(strings.Repeat("-", len(b.field)*4))
		}
	}
	fmt.Println()
}

// Check if game is over
func (b *board) checkWinCondition(sign string) bool {

	checkers := map[string]bool{
		"isDiag":   false,
		"isRow":    false,
		"isColumn": false,
	}

	// check rows
	for i, row := range b.field {
		checkers["isRow"] = true
		for j := range row {
			if b.field[i][j] != sign {
				checkers["isRow"] = false
				break
			}
		}
		if checkers["isRow"] {
			return checkers["isRow"]
		}
	}

	// check columns
	for i := range b.field {
		checkers["isColumn"] = true
		for j := range b.field {
			if b.field[j][i] != sign {
				checkers["isColumn"] = false
				break
			}
		}
		if checkers["isColumn"] {
			return checkers["isColumn"]
		}
	}

	// check diag
	checkers["isDiag"] = true
	for i := range b.field {
		if b.field[i][i] != sign {
			checkers["isDiag"] = false
			break
		}
	}

	// check reverse diag
	checkers["isDiag"] = true
	j := 0
	for i := len(b.field) - 1; i > 0; i-- {
		if b.field[i][j] != sign {
			checkers["isDiag"] = false
			break
		}
		j++
	}

	return checkers["isDiag"] || checkers["isRow"] || checkers["isColumn"]
}

func (b *board) isFilled(i, j int) bool {
	return b.field[i][j] != b.emptySign
}
