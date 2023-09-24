package route

import (
	"fmt"
	"my_interfaces/internal/vehicle"
)

type Route struct {
	Name     string
	Vehicles []vehicle.IVehicle
}

func (r *Route) AddVehicle(v vehicle.IVehicle) {
	r.Vehicles = append(r.Vehicles, v)
}

func (r *Route) PrintVehicles() {
	fmt.Println("The route vehicles:")
	for i, v := range r.Vehicles {
		fmt.Printf("%d. %v\n", i+1, v.GetName())
	}
}
