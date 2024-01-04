package main

import (
	"fmt"
	"sort"
)

var nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2} // output true
//nums = [1,2,3,4] output false
//nums = [1,2,3,1] output true

func main() {
	res := containsDuplicate(nums)
	fmt.Println(res)

	res2 := containsDuplicate2(nums)
	fmt.Println(res2)
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	for k := range m {
		if m[k] > 1 {
			return true
		}
	}

	return false
}

// solution using sorting

// Time: O(nlogn + n) = O(nlogn)
// Space: O(logn)
func containsDuplicate2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	// Time: O(nlogn)
	// Space: O(logn)
	sort.Ints(nums)

	// Time: O(n)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
