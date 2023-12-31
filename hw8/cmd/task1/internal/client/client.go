package client

import (
	"concurrency2/cmd/task1/internal/product"
	"concurrency2/cmd/task1/internal/store"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type StoreClient struct {
	Timeout time.Duration
	ctx     *context.Context

	respch   chan *product.Product
	resultch chan *Bill
}

func NewStoreClient(timeout time.Duration) *StoreClient {
	return &StoreClient{
		Timeout:  timeout,
		respch:   make(chan *product.Product),
		resultch: make(chan *Bill),
	}
}

func (c *StoreClient) HandleClient(s *store.Store) {

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*c.Timeout))
	defer cancel()

	c.ctx = &ctx
	timesToQuery := rand.Intn(3) + 2

	fmt.Printf("Client is making %d queries...\n", timesToQuery)
	for i := 0; i < timesToQuery; i++ {
		wg.Add(1)
		go RandomQueryGenerator(s, c.respch, &wg)
	}

	go func() {
		wg.Wait()
		close(c.respch)
	}()

	go PrepareBill(c.respch, c.resultch, c.ctx)

	for r := range c.resultch {
		for _, item := range r.Items {
			fmt.Printf("=====Item=====\nName: %v\nModel: %v\nManufacturer: %v\n Price: %v\n", item.Name, item.Model, item.Manufacturer, item.Price)
		}
		fmt.Printf("==============\n")
		fmt.Printf("Total price: %.2f\n", r.TotalPrice)
	}

	if err := (*c.ctx).Err(); err != nil {
		log.Fatalln(err, ":", "failed by", c.Timeout, "timeout")
	}
}
