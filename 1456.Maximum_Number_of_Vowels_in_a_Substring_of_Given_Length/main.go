package main

import "fmt"

// var s = "abciiidef"
// var k = 3

var s = "weallloveyou"
var k = 7

func main() {
	res := maxVowels(s, k)
	fmt.Println(res)
}

func maxVowels(s string, k int) int {
	vovelcountMax := 0
	counter := 0

	for i := 0; i < k; i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			counter++
		}
	}
	vovelcountMax = counter

	if k == len(s) {
		return vovelcountMax
	}

	for i := k; i < len(s); i++ {
		if s[i-k] == 'a' || s[i-k] == 'e' || s[i-k] == 'i' || s[i-k] == 'o' || s[i-k] == 'u' {
			counter--
		}
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			counter++
			vovelcountMax = max(vovelcountMax, counter)
		}
	}

	return vovelcountMax
}
