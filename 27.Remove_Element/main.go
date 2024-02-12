package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
	return p
}
