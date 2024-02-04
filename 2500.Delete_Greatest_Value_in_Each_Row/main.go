package main

import (
	"fmt"
	"sort"
)

func main() {
	grid := [][]int{{1, 2, 4}, {3, 3, 1}} //Output: 8
	res := deleteGreatestValue(grid)
	fmt.Println(res)
}

func deleteGreatestValue(grid [][]int) int {
	for _, n := range grid {
		sort.Ints(n)
	}

	sum := 0
	for i := 0; i < len(grid[0]); i++ {
		maximum := grid[0][i]
		for j := 0; j < len(grid); j++ {
			if maximum < grid[j][i] {
				maximum = grid[j][i]
			}
		}
		sum += maximum
	}

	return sum
}
