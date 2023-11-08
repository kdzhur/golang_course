package pubsub

import "fmt"

type Broker struct {
	Cons []*Consumer
	Msg  chan Message
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

type Consumer struct {
	id  string
	Msg chan Message
}

func (c *Consumer) Consume() {
	go func() {
		for {
			fmt.Println(<-c.Msg)
		}
	}()
}

type Producer struct {
	id     string
	Broker *Broker
}

func (p *Producer) Publish(msg *Message) {
	go func() {
		p.Broker.Msg <- *msg
	}()
}

type Message struct {
	Body string
}
