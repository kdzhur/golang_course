package searchtext

import "strings"

func FindStr(editorText []string, searchQuery string) []string {
	matchingLines := []string{}

	for _, line := range editorText {
		if strings.Contains(line, searchQuery) {
			matchingLines = append(matchingLines, line)
		}
	}

	return matchingLines
}
