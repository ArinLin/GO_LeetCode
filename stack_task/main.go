package main

import "fmt"

// Реализуйте структуру стека с функциональностью pop, append и print top.
func main() {
	// Создание стека
	var stack []string
	// Добавление элементов
	stack = append(stack, "world!")
	stack = append(stack, "Hello ")
	for len(stack) > 0 {
		// Print top
		n := len(stack) - 1
		fmt.Print(stack[n])
		// Pop
		stack = stack[:n]
	}
	// Output: Hello world!
}
