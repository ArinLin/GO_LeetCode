package main

import "fmt"

func main() {
	nums := []int{1, 5, 0, 3, 5} // output 3
	//nums := []int{0} // output 3
	res := minimumOperations(nums)
	fmt.Println(res)
}

func minimumOperations(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if v != 0 {
			m[v]++
		}
	}

	return len(m)
}
