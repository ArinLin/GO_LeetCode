package main

import "fmt"

var nums = []int{1, 2, 2, 4}

// Output: [2,3]

func main() {
	res := findErrorNums(nums)
	fmt.Println(res)
}

func findErrorNums(nums []int) []int {
	m := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}

	return res
}
