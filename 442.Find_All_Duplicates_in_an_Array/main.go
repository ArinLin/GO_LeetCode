package main

import "fmt"

var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}

//Output: [2,3]

func main() {
	res := findDuplicates(nums)
	fmt.Println(res)

	res2 := findDuplicates2(nums)
	fmt.Println(res2)
}

func findDuplicates(nums []int) []int {
	m := make(map[int]int)
	res := []int{}

	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}

	return res
}

func findDuplicates2(nums []int) []int {
	m := make(map[int]struct{})
	res := []int{}

	for _, num := range nums {
		if _, ok := m[num]; ok {
			res = append(res, num)
		}
		m[num] = struct{}{}
	}

	return res
}
