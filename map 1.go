package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	words := strings.Fields(text)
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Joao pedro nunes gomes da silva"

	wordCount := countWords(text)

	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
