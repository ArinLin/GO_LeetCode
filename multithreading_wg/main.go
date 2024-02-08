package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10) // мы заранее знаем, что будет выполнятся 10 горутин
	for i := 0; i < 10; i++ {
		// если не знаем сколько будет горутин, вместе wg.Add(10), можем написать wg.Add(1) строго перед тем, как запланировать выполнение горутины
		// wg.Add(1)
		go func(i int) {
			defer wg.Done() // это нужно сделать первой операцией в запущенной горутине
			// если забудем defer wg.Done(), то будет дедлок
			for j := 0; j < 10; j++ {
				fmt.Println("Hello from goroutine", i, j)
				time.Sleep(1 * time.Second)
			}
		}(i)
	}
	wg.Wait() // (ПРАВИЛЬНЫЙ ВАРИАНТ)
	//wait ждет пока в счетчике будет 0. Горутина в это время заблокирована
	//Если счетчик будет иметь отрицательное значение -- будет паника
}

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			for j := 0; j < 10; j++ {
// 				fmt.Println("Hello from goroutine", i, j)
// 				time.Sleep(1 * time.Second)
// 			}
// 		}(i)
// 	}
// 	time.Sleep(10 * time.Second) // НЕПРАВИЛЬНЫЙ ВАРИАНТ!!!
// }
