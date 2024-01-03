package main

import "fmt"

func main() {
	makeSetWithBool()
	makeSetWithStruct()
}

func makeSetWithBool() {
	set1 := map[int]bool{}

	for i := 0; i < 10; i++ {
		set1[i] = true
	}

	contains := set1[5]
	fmt.Println("set 1 contains:", contains)

	if set1[3] {
		fmt.Println("set 1 contains 3")
	}
}

// пустые структуры весят 0 байт

func makeSetWithStruct() {
	set2 := map[int]struct{}{}

	for i := 0; i < 10; i++ {
		set2[i] = struct{}{}
	}

	_, contains := set2[5]
	fmt.Println("set 2 contains:", contains)

	if _, contains = set2[3]; contains {
		fmt.Println("set 2 contains 3")
	}
}
