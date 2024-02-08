package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64 // ПРАВИЛЬНЫЙ ВАРИАНТ
	// к атомику можем обращаться с разных потоков одновременно
	// любая арифметическая операция занимает 1 шаг
	// механизм атомиков с помощью разных конструктций убеждается, что все кэши на всех ядрах сбрасываются после изменения атомика
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Add(1)
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("total count:", counter.Load()) // загружаем значение из атомика
}

// func main() {
// 	counter := 0 // НЕПРАВИЛЬНЫЙ ВАРИАНТ
// 	var wg sync.WaitGroup

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			for j := 0; j < 100; j++ {
// 				counter++
// 			}
// 		}(i)
// 	}
// 	wg.Wait()

// 	fmt.Println("total count:", counter) // НЕПРЕДСКАЗУЕМЫЙ РЕЗУЛЬТАТ
// }
