package vehicle

type Vehicle struct {
	Name  string
	Speed int
}

type Stoper interface {
	Stop()
}

type Mover interface {
	Move()
}

type NameGeter interface {
	GetName() string
}

type SpeedChanger interface {
	ChangeSpeed(speed int)
}

type IVehicle interface {
	Stoper
	Mover
	SpeedChanger
	NameGeter
}
