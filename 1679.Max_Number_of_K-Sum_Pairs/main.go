package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums := []int{1, 2, 3, 4}
	nums := []int{3, 1, 3, 4, 3}
	// k := 5 //Output: 2
	k := 6 //Output: 1
	res := maxOperations(nums, k)
	fmt.Println(res)
}

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	min := 0
	max := len(nums) - 1

	counter := 0
	for min < max {
		sum := nums[min] + nums[max]
		if sum == k {
			counter++
			min++
			max--
		} else if sum < k {
			min++
		} else {
			max--
		}
	}

	return counter
}
