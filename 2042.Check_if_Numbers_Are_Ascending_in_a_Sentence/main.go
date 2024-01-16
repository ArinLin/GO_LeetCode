package main

import (
	"fmt"
	"strconv"
	"strings"
)

var s = "1 box has 3 blue 4 red 6 green and 12 yellow marbles" // Output: true

func main() {
	res := areNumbersAscending(s)
	fmt.Println(res)
	res2 := areNumbersAscending2(s)
	fmt.Println(res2)
}

func areNumbersAscending(s string) bool {

	res := strings.Split(s, " ")

	fmt.Println(res)

	nums := []int{}

	for _, num := range res {
		isNum, err := strconv.Atoi(num)
		if err == nil {
			nums = append(nums, isNum)
		}
	}

	fmt.Println(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}

	return true
}

// second solution

func areNumbersAscending2(s string) bool {

	res := strings.Split(s, " ")

	fmt.Println(res)

	prev := -1

	for _, ch := range res {
		num, err := strconv.Atoi(ch)
		if err == nil {
			if num > prev {
				prev = num
			} else {
				return false
			}
		}
	}

	return true
}
