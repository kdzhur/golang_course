// Розробити програму, яка приймає на вхід текстові файли та здійснює різні операції з текстом, як-от підрахунок слів, заміна певних символів, видалення зайвих пробілів тощо.

// Шаблон проєктування «Стратегія» має бути використаний для реалізації різних алгоритмів обробки тексту, які можуть бути легко замінені один на одного. Наприклад, ви можете реалізувати алгоритм підрахунку слів з використанням регулярних висловів або алгоритм, який знаходить найбільше повторювані слова.

// Шаблон проєктування «Декоратор» має бути використаний для застосування різних операцій до тексту, як-от заміна символів та видалення зайвих пробілів. Наприклад, ви можете реалізувати Декоратор, який видаляє всі зайві пробіли в тексті.

// Програма має мати консольний інтерфейс користувача, який дозволяє вибирати алгоритми обробки та накладати декоратори на текст. Крім того, програма повинна виводити результати обробки тексту на екран або зберігати їх у файл. Наприклад, наприкінці програмі можна дати завдання:

// • видали (це Стратегія) із тексту всі подвійні пробіли (це Декоратор) і html теги (це теж Декоратор);

package main

import (
	"flag"
	"fmt"
	"hw_patterns/internal/editor"
	"hw_patterns/utils/filehandler"
	"log"
	"os"
)

func main() {
	action := flag.String("action", "duplications", "sets an action to apply on file's content. Possible values are: duplications, longest, title")
	output := flag.String("output", "stdout", "sets an output. Possible values are: file, stdout")
	foPath := flag.String("f", "output.txt", "Path of file to write an output to")
	flag.Parse()

	file, err := os.OpenFile("assets/file.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f := editor.NewFile(file)

	switch *action {
	case "duplications":
		f.SetAction(editor.NewDuplication())
	case "longest":
		f.SetAction(editor.NewLongestWord())
	case "title":
		f.SetAction(editor.NewTitleWord())
	default:
		log.Fatalln("Invalid input: flag action: run -h to get help.")
	}

	switch *output {
	case "stdout":
		fmt.Println(string(f.ApplyAction(editor.SetTitle, editor.ToUpper)))
	case "file":
		filehandler.CreateAndWriteToFile(f.ApplyAction(editor.SetTitle), *foPath)
	default:
		log.Fatalln("Invalid input: flag output: run -h to get help.")
	}

}
