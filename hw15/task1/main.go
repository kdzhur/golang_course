// Створіть програму, яка використовує патерн pub-sub для моніторингу системи файлів.
// Програма має стежити за змінами файлів у певному каталозі та інформувати користувача про будь-які зміни,
// що відбуваються в цьому каталозі. Кожен раз, коли зміна відбувається, програма повинна надсилати повідомлення всім підписникам,
// які підписалися на отримання сповіщень про ці зміни.

package main

import (
	"fmt"
	filesystemwatcher "qeueu/task1/internal/fileSystemWatcher"
)

func main() {
	// broker := &pubsub.Broker{
	// 	Msg: make(chan pubsub.Message),
	// }

	// broker.Accept()

	// consumer1 := &pubsub.Consumer{
	// 	Msg: make(chan pubsub.Message),
	// }
	// // fmt.Println("CONSUMER", &consumer1)
	// consumer2 := &pubsub.Consumer{
	// 	Msg: make(chan pubsub.Message),
	// }

	// producer1 := pubsub.Producer{
	// 	Broker: broker,
	// }

	// producer2 := pubsub.Producer{
	// 	Broker: broker,
	// }

	// broker.Subscribe(consumer1)
	// broker.Subscribe(consumer2)

	// producer1.Publish(&pubsub.Message{
	// 	Body: "1- TEST STRING",
	// })

	// producer2.Publish(&pubsub.Message{
	// 	Body: "2- TEST STRING",
	// })

	// consumer1.Consume()
	// consumer2.Consume()

	dirwatcher := filesystemwatcher.NewDirWatcher("D:/GOLANG/Course/hw15/task1")
	logger := filesystemwatcher.NewLogger(dirwatcher.Broker)
	logger.NotifyOnModification()

	fmt.Scanln()
	fmt.Println("Done!")

}
