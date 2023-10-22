// Реалізувати пошук слів із певним шаблоном у текстовому файлі. Задача: створити регулярний вислів,
// який можна використовувати для знаходження слів, які відповідають певному шаблону.
// Наприклад, вислів, який знаходить слова, що починаються на голосні літери та закінчуються на приголосні, або слова,
// що складаються з двох однакових букв, розділених будь-яким символом. Завдання творче.

package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text, err := os.ReadFile("assets/text_en.txt")
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	countDuplicatedWords(text)
	fmt.Println("--------")
	getTitleWords(text)
	fmt.Println("--------")
	getTheLongestWord(text)
}

func countDuplicatedWords(text []byte) {
	wordRegex := regexp.MustCompile(`\b\w+\b`)

	words := wordRegex.FindAllString(string(text), -1)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		if count > 1 {
			fmt.Printf("%s is duplicated %d times\n", word, count)
		}
	}
}

func getTitleWords(text []byte) {
	titleRegex := regexp.MustCompile(`\b[A-Z][a-z]*\b`)

	titleWords := titleRegex.FindAllString(string(text), -1)

	for _, word := range titleWords {
		fmt.Println(word)
	}
}

func getTheLongestWord(text []byte) {
	var (
		longestWord   string
		longestLength int
	)

	wordRegex := regexp.MustCompile(`\b\w+\b`)

	words := wordRegex.FindAllString(string(text), -1)

	for _, word := range words {
		if len(word) > longestLength {
			longestWord = word
			longestLength = len(word)
		}
	}
	fmt.Printf("The longest word is: %s\n", longestWord)
}
