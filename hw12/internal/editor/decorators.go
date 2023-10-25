package editor

import (
	"fmt"
	"strings"
)

func SetTitle(text string) string {
	return fmt.Sprintf("=======BEGINNING OF OUTPUT=======\n%s\n=======END OF OUTPUT=======", text)
}

func ToUpper(text string) string {
	return strings.ToUpper(text)
}
