package observer

type Observer interface {
	Update(message string)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}
