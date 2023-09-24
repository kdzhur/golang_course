package vehicle

import "fmt"

type Airplane struct {
	Vehicle
	Wingspan int
}

func (a *Airplane) Stop() {
	a.Speed = 0
	fmt.Printf("The %v is stoped. The current speed is %v km/h\n", a.Name, a.Speed)
}

func (a *Airplane) Move() {
	if a.Speed != 0 {
		fmt.Printf("The %v is moving with the %v km/h speed\n", a.Name, a.Speed)
	} else {
		fmt.Println("Trying to move. The speed is 0. The airplane is not moving")
	}
}

func (a *Airplane) ChangeSpeed(speed int) {
	a.Speed = speed
	fmt.Printf("The %v is changing speed to %v km/h\n", a.Name, a.Speed)
}

func (a *Airplane) GetName() string {
	return a.Name
}

func (a *Airplane) TakePassenger(passengers int) {
	fmt.Printf("The %v is taking %d passengers\n", a.Name, passengers)
}

func (a *Airplane) DropPassenger(passengers int) {
	fmt.Printf("The %v is droping %d passengers\n", a.Name, passengers)
}
