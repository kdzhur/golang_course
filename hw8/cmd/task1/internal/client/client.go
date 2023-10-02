package client

import (
	"concurrency2/cmd/task1/internal/product"
	"concurrency2/cmd/task1/internal/store"
	"fmt"
	"math/rand"
)

type Bill struct {
	Items      []*product.Product
	TotalPrice float64
}

func RandomQueryGenerator(s *store.Store, respch chan<- *product.Product) {
	timesToQuery := rand.Intn(3) + 2

	go func() {
		for i := 0; i < timesToQuery; i++ {
			respch <- s.Products[rand.Intn(len(s.Products))]
		}
		close(respch)
	}()

	fmt.Printf("Client is making %d queries...\n", timesToQuery)
}

func PrepareBill(respch <-chan *product.Product, resultch chan<- *Bill) {

	var total float64
	var products []*product.Product

	go func() {
		for value := range respch {
			total += value.Price
			products = append(products, value)
		}

		resultch <- &Bill{
			Items:      products,
			TotalPrice: total,
		}
		close(resultch)
	}()
}
