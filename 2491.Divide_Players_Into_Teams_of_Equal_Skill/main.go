package main

import (
	"fmt"
	"sort"
)

func main() {
	skill := []int{3, 2, 5, 1, 3, 4}
	res := dividePlayers(skill)
	fmt.Println(res)
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)

	mid := (len(skill) / 2)
	min := 0
	max := len(skill) - 1
	res := 0
	sum := skill[min] + skill[max]

	for min <= mid && max >= mid {
		if skill[min]+skill[max] != sum {
			return -1
		}
		res = res + (skill[min] * skill[max])
		min++
		max--
	}
	return int64(res)

	// res := 0
	// l := len(skill)
	// sum := skill[0] + skill[l-1]

	// for i := 0; i <= (l/2)-1; i++ {
	// 	if skill[i]+skill[l-i-1] != sum {
	// 		return int64(-1)
	// 	}
	// 	res += skill[i] * skill[l-i-1]
	// }

	// return int64(res)
}
