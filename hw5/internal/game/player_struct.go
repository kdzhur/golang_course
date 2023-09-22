package game

type player struct {
	name string
	sign string
}

// A constructor for player{}
func newPlayer(name string, sign string) *player {
	return &player{
		name: name,
		sign: sign,
	}
}
