package game

import "fmt"

func (p *player) scanPlayerMove(b *board) [2]int {
	var coordinates [2]int
	var input int

TRY_AGAIN:
	for i := range coordinates {
		fmt.Printf("%v's turn. Enter the %v position for your turn: \n", p.name, i+1)
		_, err := fmt.Scanf("%v", &input)
		if err != nil {
			fmt.Printf("invalid input: %v\n", err)
			goto TRY_AGAIN
		}
		if input < 0 || input > 2 {
			fmt.Println("invalid input: a position should be between 0 and 2")
			goto TRY_AGAIN
		}
		coordinates[i] = input
	}
	if b.isFilled(coordinates[0], coordinates[1]) {
		fmt.Println("The cell is already taken. Try again")
		goto TRY_AGAIN
	}
	return coordinates
}

func (p *player) modifyBoard(b *board, coordinates [2]int) {
	b.field[coordinates[0]][coordinates[1]] = p.sign
}
