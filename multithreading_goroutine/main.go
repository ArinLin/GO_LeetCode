package main

import (
	"fmt"
	// "runtime"
	"time"
)

func main() {
	fireGoroutine()

	// fmt.Println("after fireGoroutine")
	// fmt.Println("all my Goroutines:")
	// for t :=
	// 	time.Now(); time.Since(t) < time.Second*10; time.Sleep(time.Second) {
	// 	fmt.Println(runtime.NumGoroutine()) // выводит количество горутин
	// }
}

func fireGoroutine() {
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("Hello, hello")
	}()

	for t :=
		time.Now(); time.Since(t) < time.Second*10; time.Sleep(time.Second) {
		fmt.Println("hey")
	}
}
