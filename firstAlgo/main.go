package main

import (
	"fmt"
	"strings"
)

func main() {
	cartesianProduct()
	cartesianProduct2()
	handshakes()
	halfPyramid()
	upsideDownPyramid()
	dashPyramid()
	dashPyramidPythonStyle()
	pyramid()
	verticalPyramid()
}

func cartesianProduct() {
	fmt.Println("cartesian product")

	/*
	   +-----------------------+
	   |  X  |  A  |  B  |  C  |
	   +-----------------------+
	   |  A  | A*A | A*B | A*C |
	   |  B  | B*A | B*B | B*C |
	   |  C  | C*A | C*B | C*C |
	   +-----------------------+
	*/

	vector := [...]string{"A", "B", "C"}
	for _, x := range vector {
		for _, y := range vector {
			fmt.Printf("%s * %s\n", x, y)
		}
	}
}

func cartesianProduct2() {
	fmt.Println("cartesian product")

	/*
	   +-----------------------+
	   |  X  |  A  |  B  |  C  |
	   +-----------------------+
	   |  A  | --- | A*B | A*C |
	   |  B  | B*A | --- | B*C |
	   |  C  | C*A | C*B | --- |
	   +-----------------------+
	*/

	teams := [...]string{"Milan", "Inter", "Juventus", "Roma"}
	for _, home := range teams {
		for _, away := range teams {
			if home == away {
				continue
			}
			fmt.Printf("Home: %s, Away: %s\n", home, away)
		}
	}
}

func handshakes() {
	fmt.Println("handshakes")

	/*
	   +-----------------------------------+
	   |     |  A  |  B  |  C  |  D  |  E  |
	   +-----------------------------------+
	   |  A  | --- | A*B | A*C | A*D | A*E |
	   |  B  | --- | --- | B*C | B*D | B*E |
	   |  C  | --- | --- | --- | C*D | C*E |
	   |  D  | --- | --- | --- | --- | D*E |
	   |  E  | --- | --- | --- | --- | --- |
	   +-----------------------------------+
	*/

	people := [...]string{"Anna", "Bob", "Charlie", "Diana", "Eve"}
	var count int
	// for i, person1 := range people {
	// 	for _, person2 := range people[i+1:] {
	// 		fmt.Printf("%s X %s\n", person1, person2)
	// 		count++
	// 	}
	// }
	// fmt.Printf("Total handshakes: %d\n", count)

	// alternative solution
	for i := 0; i < len(people); i++ {
		for j := i + 1; j < len(people); j++ {
			fmt.Printf("%s X %s\n", people[i], people[j])
			count++
		}
	}
}

func halfPyramid() {
	// Example: height=4
	// *
	// **
	// ***
	// ****

	height := 6
	// for i := 0; i < height; i++ {
	// 	// How to print i+1 stars in a row?
	// 	for j := 0; j < i+1; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// alternative solution
	for i := 0; i < height; i++ {
		stars := strings.Repeat("*", i+1)
		fmt.Println(stars)
	}
}

func upsideDownPyramid() {
	// *****
	// ****
	// ***
	// **
	// *

	height := 6
	for i := height; i > 0; i-- {
		// How to print i-1 stars in a row?
		for j := i - 1; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func dashPyramid() {
	// Example: height=4
	// *---
	// **--
	// ***-
	// ****

	height := 6
	for i := 1; i <= height; i++ {
		// Print i stars followed by (height-i) dashes
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		for k := 1; k <= height-i; k++ {
			fmt.Print("-")
		}
		fmt.Println()
	}
}

func dashPyramidPythonStyle() {
	// Example: height=4
	// ----
	// *---
	// **--
	// ***-
	// ****

	height := 6

	// Print i stars followed by (height-i) dashes
	for j, k := 0, height; j <= height; j, k = j+1, k-1 {
		fmt.Print(strings.Repeat("*", j))
		fmt.Print(strings.Repeat("-", k))

		fmt.Println()
	}
}

func pyramid() {
	// Example: height=4
	//    *
	//   ***
	//  *****
	// *******

	height := 10
	// Question: how many rows do we need?
	// Question: how many stars in i-th row?
	// Question: how many spaces in i-th row?

	// Spaces:
	// k is -1 (decreasing by 1 each row) so it's spaces = constant - i
	// Get the first row (i=0) as and example: 3 = constant - 0
	// constant = 3
	// constant depends on the height of the pyramid
	// so constant = height - 1
	// Spaces formula: height - 1 - i

	// Repeat 'height' times
	for i := 0; i < height; i++ {
		// Print n-i spaces
		spaces := height - 1 - i
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}

		// Print 2*i+1 stars
		stars := 2*i + 1
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func verticalPyramid() {
	// Example
	// *
	// **
	// ***
	// ****
	// ***
	// **
	// *

	height := 7
	rows := 2*height - 1
	for i := 0; i < rows; i++ {
		// Print stars
		stars := height - abs(i-height+1)
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
