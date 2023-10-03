// Написати програму, яка симулює роботу інтернет-магазину.
// Програма має мати дві горутини: одна генерує випадкові запити від покупців (ім'я, товар, кількість),
// а друга оброблює запити й підраховує загальну вартість кожного замовлення.
// Використовуйте канали для передачі даних між горутинами та контекст для безпечної роботи з горутинами.
// Для реалізації інтерфейсу командного рядка використовуйте бібліотеку "flag".

package main

import (
	"concurrency2/cmd/task1/internal/client"
	"concurrency2/cmd/task1/internal/store"
	"flag"
	"time"
)

func main() {

	timeoutFlag := flag.Int("timeout", 500, "Set query timeout in ms. Defaults to 500ms")

	flag.Parse()

	rozetka := store.NewStore("Rozetka")
	rozetka.FillStore("assets/store.json")

	client := client.NewStoreClient(time.Duration(*timeoutFlag))

	client.HandleClient(rozetka)

	// rozetka.ShowStore()

}
