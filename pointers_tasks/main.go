package main

import (
	"fmt"
)

func main() {
	// first Task
	var numbers []*int
	for _, value := range []int{1, 2, 3, 4} {
		numbers = append(numbers, &value)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number) // Output 4 4 4 4
	}

	// second Task
	x := 3
	y := &x
	fmt.Print(*y) // Output 3
	*y = 4
	fmt.Println(x) // Output 4
}
