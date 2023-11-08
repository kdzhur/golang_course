package pubsub

import "fmt"

type Broker struct {
	Cons []*Consumer
	Msg  chan Message
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
			// fmt.Println("accept", message)
			for i := range b.Cons {
				b.Cons[i].Msg <- message
			}
		}
	}()
}

func (b *Broker) Subscribe(c *Consumer) {
	b.Cons = append(b.Cons, c)
}

type Consumer struct {
	Msg chan Message
}

func NewConsumer() *Consumer {
	return &Consumer{
		Msg: make(chan Message),
	}
}

func (c *Consumer) Consume() {
	go func() {
		for {
			fmt.Println(<-c.Msg)
		}
	}()
}

type Producer struct {
	Broker *Broker
}

func (p *Producer) Publish(msg *Message) {
	go func() {
		// fmt.Println("pub", msg)
		p.Broker.Msg <- *msg
	}()
}

type Message struct {
	Body string
}
