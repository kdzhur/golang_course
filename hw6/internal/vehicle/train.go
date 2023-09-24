package vehicle

import "fmt"

type Train struct {
	Vehicle
	wagons []string
}

func (t *Train) Stop() {
	t.Speed = 0
	fmt.Printf("The %v is stoped. The current speed is %v km/h\n", t.Name, t.Speed)
}

func (t *Train) Move() {
	if t.Speed != 0 {
		fmt.Printf("The %v is moving with the %v km/h speed\n", t.Name, t.Speed)
	} else {
		fmt.Println("Trying to move. The speed is 0. The train is not moving")
	}
}

func (t *Train) ChangeSpeed(speed int) {
	t.Speed = speed
	fmt.Printf("The %v is changing speed to %v km/h\n", t.Name, t.Speed)
}

func (t *Train) GetName() string {
	return t.Name
}

func (t *Train) TakePassenger(passengers int) {
	fmt.Printf("The %v is taking %d passengers\n", t.Name, passengers)
}

func (t *Train) DropPassenger(passengers int) {
	fmt.Printf("The %v is droping %d passengers\n", t.Name, passengers)
}
