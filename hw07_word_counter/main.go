package main

import (
	"fmt"
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

func main() {
	// Place your code here.
	// 1. написать функцию принимающую строку и возвращающую ее буз лишних знаков,
	// в нижнем регистре
	// 2. написать мапу попеременно включающую эти слова как ключи
	// сразу проверять количество этих слов в данном тексте и записывать результат как значение  в мапу
	text := "привет код"
	fmt.Println(countWords(text))
}
