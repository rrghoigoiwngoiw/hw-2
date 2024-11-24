package countwords

import (
	"strings"
	"unicode"
)

func countWords(text string) map[string]int {
	text = strings.ToLower(text)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
