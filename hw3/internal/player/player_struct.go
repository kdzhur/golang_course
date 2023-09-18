package player

type Player struct {
	Name string
}

// Player struct constructor
func NewPlayer(playerName string) *Player {
	p := &Player{
		Name: playerName,
	}

	return p
}
