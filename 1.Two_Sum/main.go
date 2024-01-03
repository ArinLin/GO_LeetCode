package main

import "fmt"

var (
	nums   = []int{1, 4, 6, 5, 12, 8}
	target = 10
)

func main() {
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, num := range nums {
		compliment := target - num
		if j, ok := m[compliment]; ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return nil
}
