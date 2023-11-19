package pubsub

type Broker struct {
	Cons          map[int]*Consumer
	Msg           chan Message
	ConsumerCount int
}

type Message struct {
	Body string
}

func NewBroker() *Broker {
	return &Broker{
		Cons: make(map[int]*Consumer),
		Msg:  make(chan Message),
	}
}

func (b *Broker) Accept() {
	go func() {
		for {
			message := <-b.Msg
			for _, c := range b.Cons {
				go func(c *Consumer) {
					b.Cons[c.ID].Msg <- message
				}(c)
			}
		}
	}()
}

func (b *Broker) Subscribe(c *Consumer) {
	b.Cons[b.ConsumerCount] = c
	c.ID = b.ConsumerCount
	b.ConsumerCount++
}
