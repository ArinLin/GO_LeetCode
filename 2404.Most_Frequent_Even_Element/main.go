package main

import "fmt"

var nums = []int{1, 2, 2, 4, 4, 1, 4} // output 2
//var nums2 = [29,47,21,41,13,37,25,7] // output -1

func main() {
	res := mostFrequentEven(nums)
	fmt.Println(res)
}

func mostFrequentEven(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		if v%2 == 0 {
			m[v]++
		}
	}

	maxFreq := 0
	res := -1

	for k, v := range m {
		if v > maxFreq || (v == maxFreq && k < res) {
			maxFreq = v
			res = k
		}

	}

	return res
}
