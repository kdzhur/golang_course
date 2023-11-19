// Розробіть програму, яка використовує патерн observer для створення ігрової системи.
// Програма має дозволяти користувачам створювати кімнати гри та запрошувати інших гравців.
// Кожен раз, коли гравець робить хід, програма повинна надсилати повідомлення всім гравцям, які підписалися на отримання сповіщень про зміни в грі.

package main

import (
	"qeueu/task2/internal/game"
)

func main() {

	player1 := game.NewPlayer("John")
	player2 := game.NewPlayer("Doe")
	players := []*game.Player{player1, player2}

	room := game.NewGameRoom("Game room", players)

	player1.TakeAnAction(room, "MOVING FORWARD")
	player2.TakeAnAction(room, "MOVING BACKWARD")

	player2.LeaveRoom(room)

	player1.TakeAnAction(room, "saying: \"I won't play alone!\"")
	player1.LeaveRoom(room)

	player1.TakeAnAction(room, "MOVING FORWARD") // no action
}
