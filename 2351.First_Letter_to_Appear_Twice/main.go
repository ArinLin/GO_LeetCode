package main

import "fmt"

var s = "adbccbaadczd" // output c

func main() {
	res := repeatedCharacter(s)
	fmt.Println(string(res))

	res2 := repeatedCharacter2(s)
	fmt.Println(string(res2))
}

// solution using struct
func repeatedCharacter(s string) byte {
	set := make(map[byte]struct{})
	for i := 0; i < len(s); i++ {
		if _, ok := set[s[i]]; ok {
			return s[i]
		}
		set[s[i]] = struct{}{}
	}
	return ' '
}

// solution using map

func repeatedCharacter2(s string) byte {
	str := []byte(s)
	m := make(map[byte]int)

	for i := 0; i < len(str); i++ {
		m[str[i]] += 1
		if m[str[i]] == 2 {
			return str[i]
		}
	}
	return 0
}
