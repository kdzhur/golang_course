// Написати програму, яка симулює роботу інтернет-магазину.
// Програма має мати дві горутини: одна генерує випадкові запити від покупців (ім'я, товар, кількість),
// а друга оброблює запити й підраховує загальну вартість кожного замовлення.
// Використовуйте канали для передачі даних між горутинами та контекст для безпечної роботи з горутинами.
// Для реалізації інтерфейсу командного рядка використовуйте бібліотеку "flag".

package main

import (
	"concurrency2/cmd/task1/internal/client"
	"concurrency2/cmd/task1/internal/product"
	"concurrency2/cmd/task1/internal/store"
	"context"
	"fmt"
	"time"
)

func main() {

	respch := make(chan *product.Product)
	resultch := make(chan *client.Bill)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
	defer cancel()

	rozetka := store.NewStore("Rozetka")

	rozetka.FillStore("assets/store.json")

	// rozetka.ShowStore()

	go client.RandomQueryGenerator(rozetka, respch)
	go client.PrepareBill(respch, resultch, ctx)

	for r := range resultch {
		for _, item := range r.Items {
			fmt.Printf("=====Item=====\nName: %v\nModel: %v\nManufacturer: %v\n Price: %v\n", item.Name, item.Model, item.Manufacturer, item.Price)
		}
		fmt.Printf("==============\n")
		fmt.Printf("Total price: %.2f\n", r.TotalPrice)
	}
}
