package editor

import (
	"fmt"
	"regexp"
)

type Duplication struct{}

func (d *Duplication) WorkOnText(f *File) []byte {
	var tmp string

	wordRegex := regexp.MustCompile(`\b\w+\b`)

	words := wordRegex.FindAllString(string(f.Content), -1)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	for word, count := range wordCount {
		if count > 1 {
			tmp += fmt.Sprintf("%s is duplicated %d times\n", word, count)
		}
	}

	return []byte(tmp)
}

func NewDuplication() *Duplication {
	return &Duplication{}
}
