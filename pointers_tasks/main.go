package main

import (
	"fmt"
)

func main() {
	var numbers []*int
	for _, value := range []int{1, 2, 3, 4} {
		numbers = append(numbers, &value)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number) // Output 4 4 4 4
	}
}
