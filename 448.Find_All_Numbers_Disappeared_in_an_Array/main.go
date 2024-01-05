package main

import "fmt"

var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}

// Output: [5,6]

func main() {
	res := findDisappearedNumbers(nums)
	fmt.Println(res)
}

func findDisappearedNumbers(nums []int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}
