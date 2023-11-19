package observer

type Observer interface {
	Update(message string)
}

type Subject interface {
	registerObserver(observer Observer)
	removeObserver(observer Observer)
	notifyObservers(message string)
}
