package main

import "fmt"

var s = "foobar"
var letter byte = 'o'

//Output: 33

func main() {
	res := percentageLetter(s, letter)
	fmt.Println(res)

	res2 := percentageLetter2(s, letter)
	fmt.Println(res2)
}

func percentageLetter(s string, letter byte) int {
	counter := 0

	ss := []byte(s)

	for _, ch := range ss {
		if ch == letter {
			counter++
		}
	}

	res := 0
	res = int((counter * 100) / len(s))
	return res
}

func percentageLetter2(s string, letter byte) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return (count * 100) / len(s)
}
