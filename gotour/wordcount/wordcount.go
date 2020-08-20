package wordcount

import "strings"

func wordCount(s string) map[string]int {
	wordCount := map[string]int{}

	for _, word := range strings.Fields(s) {
		wordCount[word]++
	}

	return wordCount
}
