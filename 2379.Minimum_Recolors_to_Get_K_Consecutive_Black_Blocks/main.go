package main

import "fmt"

func main() {
	// blocks := "WBBWWBBWBW"
	// k := 7
	blocks := "BWWBWWWWBWWWWBBWWW"
	k := 7
	res := minimumRecolors(blocks, k)
	fmt.Println(res)
}

func minimumRecolors(blocks string, k int) int {
	whites := 0

	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			whites++
		}
	}

	if whites == 0 {
		return 0
	}

	minR := whites

	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whites++
		}

		if blocks[i-k] == 'W' {
			whites--
		}

		if whites < minR {
			minR = whites
		}
	}

	return minR
}
