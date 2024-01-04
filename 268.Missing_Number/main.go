package main

import "fmt"

var nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1} // output 8

func main() {
	res := missingNumber(nums)
	fmt.Println(res)

	res2 := missingNumber2(nums)
	fmt.Println(res2)
}

// using struct
func missingNumber(nums []int) int {
	m := make(map[int]struct{})

	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, exsist := m[i]; !exsist {
			return i
		}
	}

	return 0
}

// using map int int
func missingNumber2(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	for i := 0; i <= len(nums); i++ {
		if m[i] == 0 {
			return i
		}
	}

	return 0
}
