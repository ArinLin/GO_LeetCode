package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello, my name is John"
	res := countSegments(s)
	fmt.Println(res)
	res2 := countSegments2(s)
	fmt.Println(res2)
}

func countSegments(s string) int {
	if len(s) == 0 {
		return 0
	}
	counter := 0
	flag := false
	for _, ch := range s {
		if !flag && ch != ' ' {
			flag = true
			counter++
		}
		if flag && ch == ' ' {
			flag = false
		}
	}
	return counter
}

// another solution
func countSegments2(s string) int {
	fmt.Println(strings.Fields(s))
	return len(strings.Fields(s))
}
