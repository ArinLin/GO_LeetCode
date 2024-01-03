package main

import "fmt"

var nums = []int{2, 11, 2, 1, 1, 1, 2, 11, 2, 3, 3, 11, 11, 11, 3}

func main() {
	ans := majorityElement(nums)
	fmt.Println(ans)
}

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	res := 0
	max := 0
	for i, num := range m {
		if num > max {
			max = num
			res = i
		}
	}

	return res
}
