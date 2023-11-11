// Розробіть програму, яка використовує патерн observer для створення ігрової системи.
// Програма має дозволяти користувачам створювати кімнати гри та запрошувати інших гравців.
// Кожен раз, коли гравець робить хід, програма повинна надсилати повідомлення всім гравцям, які підписалися на отримання сповіщень про зміни в грі.

package main

import "qeueu/task2/internal/game"

func main() {
	room := game.NewGameRoom("Game room")

	player1 := game.NewPlayer("John")
	player2 := game.NewPlayer("Doe")

	room.RegisterObserver(player1)
	room.RegisterObserver(player2)

	room.NotifyObservers("Player Bob made a move.")

	room.RemoveObserver(player1)
	room.NotifyObservers("Player Alice made a move.")
}
