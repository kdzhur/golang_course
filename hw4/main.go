package main

import (
	"arraysAndSlices/internal/averagegrade"
	"arraysAndSlices/internal/searchtext"
	"bufio"
	"fmt"
	"os"
)

func main() {

	editorText := []string{}
	grades := []float64{90.5, 85.0, 78.5, 92.0, 88.5}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter the text (enter 'q' to finish): ")
		scanner.Scan()
		input := scanner.Text()

		if input == "q" {
			break
		}

		editorText = append(editorText, input)
	}

	fmt.Println("Enter text to find: ")
	scanner.Scan()
	searchQuery := scanner.Text()

	matchingLines := searchtext.FindStr(editorText, searchQuery)

	fmt.Println("The searching results:")
	if len(matchingLines) == 0 {
		fmt.Println("NONE")
	} else {
		for _, line := range matchingLines {
			fmt.Println(line)
		}
	}

	averageGrade := averagegrade.GetAverageGrade(grades)

	fmt.Printf("Average grade: %.2f\n", averageGrade)
}
