package pubsub

import "fmt"

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
			msg := <-c.Msg
			fmt.Println(msg.Body)
		}
	}()
}
