package main

import "fmt"

var s = "()[]{}" // Output: true

func main() {
	res := isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {
	stack := []byte{}
	ss := []byte(s)

	for _, ch := range ss {
		switch ch {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default:
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
