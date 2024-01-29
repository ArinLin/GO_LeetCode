package main

import "fmt"

var sentence = "thequickbrownfoxjumpsoverthelazydog" //Output: true
//var sentence = "leetcode" //Output: false

func main() {
	res := checkIfPangram(sentence)
	fmt.Println(res)
}

func checkIfPangram(sentence string) bool {
	a := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

	m := make(map[rune]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range sentence {
		m[v]++
	}

	for _, v := range m {
		if v == 1 {
			return false
		}
	}

	return true
}
