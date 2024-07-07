package main

import (
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	words := [][]string{}
	wordsMap := map[string][]string{}

	for _, word := range strs {
		wordsMap[sortize(word)] = append(wordsMap[sortize(word)], word)
	}

	for _, wordsSlice := range wordsMap {
		words = append(words, wordsSlice)
	}

	return words
}

func sortize(word string) string {

	wordsSlice := strings.Split(word, "")

	for i := 0; i < len(wordsSlice); i++ {
		for k := 0; k < len(wordsSlice); k++ {
			if wordsSlice[i] < wordsSlice[k] {
				temp := wordsSlice[i]
				wordsSlice[i] = wordsSlice[k]
				wordsSlice[k] = temp
			}
		}
	}

	return strings.Join(wordsSlice, "")
}

func main() {
}
