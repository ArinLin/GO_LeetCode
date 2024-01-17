package main

import "fmt"

var word1 = "abcdeef"
var word2 = "abaaacc"

func main() {
	res := checkAlmostEquivalent(word1, word2)
	fmt.Println(res)
}

func checkAlmostEquivalent(word1 string, word2 string) bool {
	m1 := make(map[rune]int)
	allowed := 3

	for _, ch := range word1 {
		m1[ch]++
	}

	for _, ch := range word2 {
		m1[ch]--
	}

	for _, val := range m1 {
		if val > allowed || val < -allowed {
			return false
		}
	}

	return true
}
