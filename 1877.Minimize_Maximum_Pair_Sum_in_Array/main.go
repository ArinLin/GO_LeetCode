package main

import (
	"fmt"
	"sort"
)

var nums = []int{3, 5, 4, 2, 4, 6} // Output: 8

func main() {
	res := minPairSum(nums)
	fmt.Println(res)
}

func minPairSum(nums []int) int {
	var l = 0
	var r = len(nums) - 1
	var maximum = 0

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for l < r {
		maximum = max(maximum, nums[l]+nums[r])
		l += 1
		r -= 1
	}

	return maximum
}
