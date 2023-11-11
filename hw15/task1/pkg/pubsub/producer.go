package pubsub

type Producer struct {
	Broker *Broker
}

func (p *Producer) Publish(msg *Message) {
	go func() {
		p.Broker.Msg <- *msg
	}()
}
