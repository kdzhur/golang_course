// 1. Виконати пошук телефонних номерів у файлі з даними контактів. Задача: створити регулярний вислів,
// який можна використовувати для знаходження телефонних номерів, записаних у різних форматах.
// Наприклад, ви можете почати з використання вислову, який знаходить номери телефонів, що складаються з 10 цифр,
// а потім розширити його, додавши підтримку різних форматів, наприклад, номери з круглими дужками, пробілами та дефісами.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.OpenFile("assets/numbers.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		re := regexp.MustCompile(`(?:\d{10}|\(\d{3}\)\s?\d{3}[-.]?\d{4}|\d{3}[-.]?\d{3}[-.]?\d{4})`)
		submatch := re.FindStringSubmatch(line)
		if len(submatch) == 0 {
			fmt.Println(line, "is not valid phone number")
		} else {
			fmt.Println(submatch[0], "is a valid phone number")
		}
	}
}
