package parser

import (
	"fmt"
	"strings"
)

func WordCountFromText(str string) []string {
	words := strings.Split(str, " ")
	wordsWithCount := make(map[string]int, len(words))
	for _, word := range words {
		word = strings.Replace(word, " ", "", -1)
		word = strings.Replace(word, "\n", "", -1)
		word = strings.Replace(word, "\t", "", -1)
		wordsWithCount[word]++
	}
	var wordsSlice []string
	for key, value := range wordsWithCount {
		wordsSlice = append(wordsSlice, fmt.Sprintf("%s - %d", key, value))
	}
	return wordsSlice
}
