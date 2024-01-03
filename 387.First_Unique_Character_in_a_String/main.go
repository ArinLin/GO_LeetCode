package main

import (
	"fmt"
)

var s = "loveleetcode" // output 2 (index of the first non-repeating character)

func main() {
	res := firstUniqChar(s)
	fmt.Println(res)
}

func firstUniqChar(s string) int {
	m := make(map[rune]int)

	// add 1 on the value portion of the map
	for _, ch := range s {
		m[ch]++
	}

	for i, value := range s {
		if m[value] == 1 {
			return i
		}
	}
	return -1
}
