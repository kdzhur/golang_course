package vehicle

import "fmt"

type Car struct {
	Vehicle
	Brand string
}

func (c *Car) Stop() {
	c.Speed = 0
	fmt.Printf("The %v is stoped. The current speed is %v km/h\n", c.Name, c.Speed)
}

func (c *Car) Move() {
	if c.Speed != 0 {
		fmt.Printf("The %v is moving with the %v km/h speed\n", c.Name, c.Speed)
	} else {
		fmt.Println("Trying to move. The speed is 0. A car is not moving")
	}
}

func (c *Car) ChangeSpeed(speed int) {
	c.Speed = speed
	fmt.Printf("The %v is changing speed to %v km/h\n", c.Name, c.Speed)
}

func (c *Car) GetName() string {
	return c.Name
}

func (c *Car) TakePassenger(passengers int) {
	fmt.Printf("The %v is taking %d passengers\n", c.Name, passengers)
}

func (c *Car) DropPassenger(passengers int) {
	fmt.Printf("The %v is droping %d passengers\n", c.Name, passengers)
}
