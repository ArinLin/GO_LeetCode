package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Foo struct { // 16
	w bool // 1 байт
	x bool // 1 байт
	// пустое пространство для выравнивание в 6 байт
	y uint64 // 8 байт
}
type Bar struct { // 24
	x bool // 1 байт
	// пустое пространство для выравнивание в 7 байт
	y uint64 // 8 байт
	w bool   // 1 байт
	// пустое пространство для выравнивание в 7 байт
}

func main() {
	fmt.Println(runtime.GOARCH)
	newFoo := new(Foo)
	fmt.Println(unsafe.Sizeof(*newFoo))
	newBar := new(Bar)
	fmt.Println(unsafe.Sizeof(*newBar))
}
