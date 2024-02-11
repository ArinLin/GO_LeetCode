package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)
}

func removeDuplicates(nums []int) int {
	p := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[p] = nums[i]
			p++
		}
	}

	return p
}
