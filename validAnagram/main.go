package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charsMap := map[rune]int{}
	for _, char := range s {
		charsMap[char]++
	}

	for _, char := range t {
		if charsMap[char] > 0 {
			charsMap[char]--
		} else {
			return false
		}
	}

	return true
}

func main() {
	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)

	fmt.Println(res)
}
