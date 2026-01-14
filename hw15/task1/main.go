// Створіть програму, яка використовує патерн pub-sub для моніторингу системи файлів.
// Програма має стежити за змінами файлів у певному каталозі та інформувати користувача про будь-які зміни,
// що відбуваються в цьому каталозі. Кожен раз, коли зміна відбувається, програма повинна надсилати повідомлення всім підписникам,
// які підписалися на отримання сповіщень про ці зміни.

package main

import (
	"fmt"
	"log"
	"os"
	filesystemwatcher "qeueu/task1/internal/fileSystemWatcher"
)

const RATE_LIMIT = 1

func main() {

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dirwatcher := filesystemwatcher.NewDirWatcher(workDir, RATE_LIMIT)
	logger := filesystemwatcher.NewLogger(dirwatcher.Broker)
	logger.NotifyOnModification()

	fmt.Scanln()
	fmt.Println("Done!")

}
