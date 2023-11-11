package game

import (
	"fmt"
	"qeueu/task2/pkg/observer"
)

type Player struct {
	Name string
}

type GameRoom struct {
	Name      string
	Players   []Player
	Observers []observer.Observer
}

func NewPlayer(name string) Player {
	return Player{
		Name: name,
	}
}

func NewGameRoom(name string) GameRoom {
	return GameRoom{
		Name:      name,
		Players:   make([]Player, 0),
		Observers: make([]observer.Observer, 0),
	}
}

func (gr *GameRoom) RegisterObserver(observer observer.Observer) {
	gr.Observers = append(gr.Observers, observer)
}

func (gr *GameRoom) RemoveObserver(observer observer.Observer) {
	for i, obs := range gr.Observers {
		if obs == observer {
			gr.Observers = append(gr.Observers[:i], gr.Observers[i+1:]...)
			break
		}
	}
}

func (gr *GameRoom) NotifyObservers(message string) {
	for _, observer := range gr.Observers {
		observer.Update(message)
	}
}

func (p Player) Update(message string) {
	fmt.Printf("Player %s received message: %s\n", p.Name, message)
}
