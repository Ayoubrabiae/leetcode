package main

import "fmt"

func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return []int{}
	}

	res := []int{}
	wordsMap := sliceToMap(words)
	lenOfWord := len(words[0])
	lenOfSub := len(words) * lenOfWord
	left := 0
	right := lenOfSub

	for right < len(s)+1 {
		sub := s[left:right]
		if compare(sub, &wordsMap, &lenOfWord) {
			res = append(res, left)
		}

		left++
		right++
	}

	return res
}

func sliceToMap(slc []string) map[string]int {
	res := map[string]int{}

	for _, word := range slc {
		res[word]++
	}

	return res
}

func compare(sub string, wordsMap *map[string]int, lenOfWod *int) bool {
	wordsMapCopy := mapCopier(wordsMap)
	left := 0
	right := *lenOfWod

	for right < len(sub)+1 {
		word := sub[left:right]

		if wordsMapCopy[word] > 0 {
			wordsMapCopy[word]--
		} else {
			return false
		}

		left += *lenOfWod
		right += *lenOfWod
	}

	return true
}

func mapCopier(m *map[string]int) map[string]int {
	res := map[string]int{}

	for key, value := range *m {
		res[key] = value
	}

	return res
}

func main() {
	str := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	result := findSubstring(str, words)

	fmt.Println(result)
}
