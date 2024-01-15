package main

import "fmt"

var accounts = [][]int{
	{1, 5},
	{7, 3},
	{3, 5},
}

func main() {
	res := maximumWealth(accounts)
	fmt.Println(res)
}

func maximumWealth(accounts [][]int) int {
	maximum := 0

	for _, num := range accounts {
		sum := 0
		for _, element := range num {
			sum += element
			maximum = max(maximum, sum)
		}
	}

	return maximum
}
