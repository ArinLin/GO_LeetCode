package main

import "fmt"

// var nums1 = []int{1, 2, 2, 1}
// var nums2 = []int{2, 2}

//output [2,2]

var nums1 = []int{4, 9, 5}
var nums2 = []int{9, 4, 9, 8, 4}

//output [4,9]

func main() {
	res := intersect(nums1, nums2)
	fmt.Println(res)
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for _, v := range nums1 {
		m[v]++
	}

	res := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num] -= 1
		}
	}

	return res
}
