package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []int{3, 2, 2, 1}
	limit := 3 // Output 3 boats (1, 2), (2) and (3)
	res := numRescueBoats(people, limit)
	fmt.Println(res)
}

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	min := 0
	max := len(people) - 1
	counter := 0

	for min <= max {
		if people[min]+people[max] <= limit {
			min++
		}
		max--
		counter++
	}

	return counter
}
