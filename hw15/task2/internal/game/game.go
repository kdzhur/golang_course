package game

import (
	"fmt"
	"qeueu/task2/pkg/observer"
)

type Player struct {
	name string
}

type GameRoom struct {
	Name      string
	Players   []*Player
	Observers []observer.Observer
}

func NewPlayer(name string) *Player {
	return &Player{
		name: name,
	}
}

func NewGameRoom(name string, players []*Player) *GameRoom {
	var observers []observer.Observer

	for _, player := range players {
		observers = append(observers, player)
	}

	g := &GameRoom{
		Name:      name,
		Players:   make([]*Player, 0),
		Observers: make([]observer.Observer, 0),
	}

	g.initGame(observers)

	return g
}

func (gr *GameRoom) initGame(observers []observer.Observer) {
	for _, ob := range observers {
		gr.registerObserver(ob)
	}
}

func (p *Player) TakeAnAction(gr *GameRoom, action string) {
	a := fmt.Sprintf("Player %s is %s\n", p.name, action)
	gr.notifyObservers(a)
}

func (p *Player) LeaveRoom(gr *GameRoom) {
	for i, player := range gr.Players {
		if player == p {
			gr.Players = append(gr.Players[:i], gr.Players[i+1:]...)
			break
		}
	}

	gr.removeObserver(p)
}

func (gr *GameRoom) registerObserver(observer observer.Observer) {
	gr.Observers = append(gr.Observers, observer)
}

func (gr *GameRoom) removeObserver(observer observer.Observer) {
	for i, obs := range gr.Observers {
		if obs == observer {
			gr.Observers = append(gr.Observers[:i], gr.Observers[i+1:]...)
			break
		}
	}
}

func (p *Player) Update(message string) {
	fmt.Printf("Player %s received message: %s\n", p.name, message)
}

func (gr *GameRoom) notifyObservers(message string) {
	for _, observer := range gr.Observers {
		observer.Update(message)
	}
}
