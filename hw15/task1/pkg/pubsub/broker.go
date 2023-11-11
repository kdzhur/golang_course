package pubsub

type Broker struct {
	Cons []*Consumer
	Msg  chan Message
}

type Message struct {
	Body string
}

func NewBroker() *Broker {
	return &Broker{
		Msg: make(chan Message),
	}
}

func (b *Broker) Accept() {
	go func() {
		for {
			message := <-b.Msg
			for i := range b.Cons {
				b.Cons[i].Msg <- message
			}
		}
	}()
}

func (b *Broker) Subscribe(c *Consumer) {
	b.Cons = append(b.Cons, c)
}
