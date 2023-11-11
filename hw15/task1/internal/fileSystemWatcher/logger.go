package filesystemwatcher

import "qeueu/task1/pkg/pubsub"

type Logger struct {
	cons *pubsub.Consumer
}

func NewLogger(broker *pubsub.Broker) *Logger {
	cons := pubsub.NewConsumer()
	broker.Subscribe(cons)

	return &Logger{
		cons: cons,
	}
}

func (l *Logger) NotifyOnModification() {
	l.cons.Consume()
}
