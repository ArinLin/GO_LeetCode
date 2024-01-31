package main

import (
	"fmt"
)

func main() {
	cardPoints := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3 //Output: 12
	// cardPoints := []int{9, 7, 7, 9, 7, 7, 9}
	// k := 7 //Output: 55
	res := maxScore(cardPoints, k)
	fmt.Println(res)
}

func maxScore(cardPoints []int, k int) int {
	l := k - 1
	r := len(cardPoints) - 1

	current := 0
	for _, card := range cardPoints[0:k] {
		current += card
	}

	maxPoint := current
	for i := 0; i < k; i++ {
		current += (cardPoints[r] - cardPoints[l])
		maxPoint = max(maxPoint, current)
		l--
		r--
	}

	return maxPoint
}
