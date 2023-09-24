package passenger

import "my_interfaces/internal/vehicle"

type IPassenger interface {
	vehicle.IVehicle
	TakePassenger()
	DropPassenger()
}
