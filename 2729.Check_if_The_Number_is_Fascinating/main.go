package main

import (
	"fmt"
	"strconv"
)

var n = 192 //Output: true
// var n = 100 //Output: false

func main() {
	res := isFascinating(n)
	fmt.Println(res)
}

func isFascinating(n int) bool {
	nStr := strconv.Itoa(n)
	n2Str := strconv.Itoa(n * 2)
	n3Str := strconv.Itoa(n * 3)

	concat := nStr + n2Str + n3Str

	for runeCount := range concat {
		runeCount++
		if runeCount > 9 {
			return false
		}
	}

	targetChar := '0'
	for _, char := range concat {
		if char == targetChar {
			return false
		}
	}

	m := make(map[rune]int)

	for _, v := range concat {
		m[v]++
	}

	for _, v := range m {
		if v == 0 || v > 1 {
			return false
		}
	}

	return true
}
