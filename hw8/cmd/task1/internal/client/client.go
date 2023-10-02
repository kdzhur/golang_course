package client

import (
	"concurrency2/cmd/task1/internal/product"
	"concurrency2/cmd/task1/internal/store"
	"context"
	"fmt"
	"log"
	"time"
)

type StoreClient struct {
	Timeout time.Duration
	Ctx     *context.Context
}

func NewStoreClient(timeout time.Duration) *StoreClient {
	return &StoreClient{
		Timeout: timeout,
	}
}

func (c *StoreClient) HandleClient(s *store.Store) {
	respch := make(chan *product.Product)
	resultch := make(chan *Bill)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*c.Timeout))
	defer cancel()
	c.Ctx = &ctx

	go RandomQueryGenerator(s, respch)
	go PrepareBill(respch, resultch, c.Ctx)

	for r := range resultch {
		for _, item := range r.Items {
			fmt.Printf("=====Item=====\nName: %v\nModel: %v\nManufacturer: %v\n Price: %v\n", item.Name, item.Model, item.Manufacturer, item.Price)
		}
		fmt.Printf("==============\n")
		fmt.Printf("Total price: %.2f\n", r.TotalPrice)
	}

	if err := (*c.Ctx).Err(); err != nil {
		log.Fatalln(err, ":", "failed by", c.Timeout, "ms timeout")
	}
}
