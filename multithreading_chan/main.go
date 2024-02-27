package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// var ch = make(chan int)
// var read <- chan int = ch //это канал только для чтения. То есть, можно только читать значения из этого канала, но нельзя отправлять в него значения.
// var write chan<- int = ch //это канал только для записи. Можно только отправлять значения в этот канал, но нельзя читать из него.

func main() {
	//channelRange()
	//channelDirection()
	//readingFromClosedChannel()
	//readingFromMultipleGoroutines()
	//semaphoreTest()
	//writeToContext()
	contextCancel()
}
func channelRange() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// Закрываем канал, чтобы сообщить о том, что больше данных не будет
		// Если канал не закрыть, то цикл по нему не завершится и мы получим панику из-за дедлока
		close(ch)
	}()
	for n := range ch { // Если канал не закрыть, то эта горутина будет ждать следущего элемента, из-за чего будет дедлок
		fmt.Println(n)
	}
	fmt.Println("done")
}

func channelDirection() {
	// ch - read-only, можно только читать, нельзя писать и закрывать
	ch := producer(0, 100, 10) //(start: 0, end: 100, step: 10)
	// close(ch) // Ошибка компиляции
	// ch <- 10 // Ошибка компиляции
	consumer(ch)
}

func producer(start, end, step int) <-chan int {
	// Создаем обычный канал, но возвращаем его как read-only
	// Это меняет тип только формально. На самом деле это просто подсказка компилятору
	ch := make(chan int)
	go func() {
		for cur := start; cur < end; cur += step {
			ch <- cur
		}
		close(ch)
	}()
	return ch
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Println(n)
	}
}

func readingFromClosedChannel() {
	ch := make(chan int, 1)
	ch <- 10

	n, ok := <-ch
	fmt.Println("Before closing:", n, ok)
	close(ch)
	n, ok = <-ch
	fmt.Println("After closing 1:", n, ok)
	n, ok = <-ch
	fmt.Println("After closing 2:", n, ok)
}

func readingFromMultipleGoroutines() {
	// Известен как паттерн "workers pool"
	const messages = 50
	var goroutines = 4 //runtime.NumCPU()

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func(j int) {
			defer wg.Done()

			for n := range ch {
				fmt.Printf("got %d from channel in goroutine %d\n", n, j)
			}
			fmt.Printf("goroutine %d finished\n", j)
		}(i)
	}
	for i := 0; i < messages; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
	//Вывод:
	// Если читает несколько горутин из одного и того же канала, то сообщения распределяются между ними, не дублируются
	// close сигнал получают все читатели сразу. Это заканчивает циклы и все последующие чтения будут читать пару (0, false)
}

func semaphoreTest() {
	const messages = 50
	s := semaphore.NewWeighted(4)

	var wg sync.WaitGroup

	for i := 0; i < messages; i++ {
		func() {
			err := s.Acquire(context.Background(), 1)
			if err != nil {
				log.Fatal(err)
			}

			wg.Add(1)
			go func(i int) {
				defer s.Release(1)
				defer wg.Done()
				fmt.Printf("got: %d, current nums of goroutines: %d\n", i, runtime.NumGoroutine()-1)
			}(i)
		}()
	}
	wg.Wait()
	// В отличии от предыдущего примера, будет запущено 50 горутин, но одновременно будет работать не больше 4
}

func writeToContext() {
	// a := 5
	// ctx := context.Background()
	// ctxCild := context.WithValue(ctx, "key", a)

	// value := ctxChild.Value("key")
	// if value != nil {
	// 	fmt.Println("Value in context:", value)
	// } else {
	// 	fmt.Println("Value not found in context")
	// }
}

func contextCancel() {
	// Создаем фоновый контекст и функцию cancel, которая будет отменять контекст
	ctxBack := context.Background()
	ctx, cancel := context.WithCancel(ctxBack)

	// Запускаем функцию, которая будет мониторить контекст
	go monitorContext(ctx)

	// Отменяем контекст через 3 секунды
	time.Sleep(3 * time.Second)
	cancel()

	// Ждем 1 секунду, чтобы увидеть результат отмены
	time.Sleep(1 * time.Second)
}

func monitorContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): //мы блокируем выполнение программы до тех пор, пока не будет получено сообщение из канала ctx.Done(). Это означает, что программа ожидает, пока контекст не будет отменен или закончится его срок действия, прежде чем продолжить выполнение кода дальше.
			// Если контекст отменен, выводим сообщение и завершаем функцию
			fmt.Println("Context is canceled")
			return
		default:
			// Имитация работы
			time.Sleep(1 * time.Second)
			fmt.Println("Working...")
		}
	}
}
