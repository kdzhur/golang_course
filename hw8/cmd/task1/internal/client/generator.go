package client

import (
	"concurrency2/cmd/task1/internal/product"
	"concurrency2/cmd/task1/internal/store"
	"context"
	"math/rand"
	"sync"
	"time"
)

type Bill struct {
	Items      []*product.Product
	TotalPrice float64
}

func RandomQueryGenerator(s *store.Store, respch chan<- *product.Product, wg *sync.WaitGroup) {
	respch <- s.Products[rand.Intn(len(s.Products))]
	wg.Done()
}

func PrepareBill(respch <-chan *product.Product, resultch chan<- *Bill, ctx *context.Context) {

	var total float64
	var products []*product.Product

	for {
		select {
		case v, ok := <-respch:
			// simulate latency
			time.Sleep(time.Duration(time.Millisecond) * time.Duration(rand.Intn(100)+30))
			if ok {
				total += v.Price
				products = append(products, v)
			} else {
				resultch <- &Bill{
					Items:      products,
					TotalPrice: total,
				}
				close(resultch)
				return
			}
		case <-(*ctx).Done():
			resultch <- &Bill{
				Items:      products,
				TotalPrice: total,
			}
			close(resultch)
			return
		}
	}
}
