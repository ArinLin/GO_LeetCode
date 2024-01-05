package main

import "fmt"

var nums1 = []int{1, 2, 2, 1}
var nums2 = []int{2, 2}

// output[2]

func main() {
	res := intersection(nums1, nums2)
	fmt.Println(res)
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	answer := []int{}

	for _, num := range nums1 {
		m[num] = true
	}

	for _, n := range nums2 {
		if m[n] == true {
			answer = append(answer, n)
			m[n] = false
		}
	}
	return answer
}
