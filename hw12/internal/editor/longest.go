package editor

import (
	"strings"
)

type LongestWord struct{}

func (d *LongestWord) WorkOnText(f *File) []byte {
	longest, length := "", 0
	for _, word := range strings.FieldsFunc(string(f.Content), func(r rune) bool {
		return r == '\n' || r == ' '
	}) {
		trimmedWord := strings.TrimFunc(word, func(r rune) bool {
			return r == '.' || r == ',' || r == ':' || r == ';' || r == '"' || r == '\''
		})
		if len(trimmedWord) > length {
			longest, length = trimmedWord, len(word)
		}
	}

	return []byte(longest)
}

func NewLongestWord() *LongestWord {
	return &LongestWord{}
}
