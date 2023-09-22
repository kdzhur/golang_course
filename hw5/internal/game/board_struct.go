package game

type board struct {
	field     [3][3]string
	emptySign string
}

// A constructor for board{}
func newBoard(emptySign string) *board {
	return &board{
		field: [3][3]string{
			{emptySign, emptySign, emptySign},
			{emptySign, emptySign, emptySign},
			{emptySign, emptySign, emptySign},
		},
		emptySign: emptySign,
	}
}
