package editor

import (
	"fmt"
	"regexp"
)

type TitleWord struct{}

func (d *TitleWord) WorkOnText(f *File) []byte {
	var result string

	titleRegex := regexp.MustCompile(`\b[A-Z][a-z]*\b`)
	titleWords := titleRegex.FindAllString(string(f.Content), -1)

	for _, word := range titleWords {
		result += fmt.Sprintln(word)
	}

	return []byte(result)
}

func NewTitleWord() *TitleWord {
	return &TitleWord{}
}
