package main

import (
	"fmt"
	"sync"
)

// Запись в мапу
func main() {
	m := make(map[int]int)
	var mx sync.Mutex // добавляем мютекс
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mx.Lock()    // Блокируем
				m[i] = i * j // Эта часть кода называется критической секцией
				mx.Unlock()  // Разблокируем
				// все записи в мапу будут происходить друг за другом
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Lenght", len(m))
}

// или критическую секцию записывают через анономную функция -- это тоже ПРАВИЛЬНО

func WriteInMap() {
	m := make(map[int]int)
	var mx sync.Mutex // добавляем мютекс
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				func() { // Записываем секцию в анономную функцию, чтобы отработал defer mx.Unlock()
					mx.Lock()
					defer mx.Unlock()
					m[i] = i * j // Эта часть кода называется критической секцией
				}()
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Lenght", len(m))
}

// чтение из мапы
func readFromMap() {
	m := make(map[int]int)
	var mx sync.RWMutex // добавляем рвмютекс |1 writer OR n readers|
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				func() {
					mx.RLock() // пока кто-то читает, любой может присоединиться и тоже читать, но никто не может взять блокировку на запись
					defer mx.RUnlock()
					_ = m[i] // Эта часть кода называется критической секцией
				}()
			}
		}(i)
	}
	wg.Wait()

	fmt.Println("Lenght", len(m))
}

// func main() {
// 	m := make(map[int]int) // НЕПРАВИЛЬНЫЙ ВАРИАНТ
// 	var wg sync.WaitGroup

// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			for j := 0; j < 100; j++ {
// 				m[i] = i * j
// 			}
// 		}(i)
// 	}
// 	wg.Wait()

// 	fmt.Println("Lenght", len(m)) // Будет паника
// }
