package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int)

	go (func() {
		for i := 0; i < 100; i++ {
			numbers <- i
		}

		close(numbers)
	})()

	parallelSquarer(numbers)
}

func parallelSquarer(numbers <-chan int) {
	type result struct {
		value int
		err   error
	}

	results := make(chan result)

	var wg sync.WaitGroup // ( !!! ) Для выхода из функции после завершения всех горутин
	for n := range numbers {
		wg.Add(1) // Увеличим счетчик потоков
		go func(number int) {
			defer wg.Done()
			// Эмуляция ошибки при выполнении горутины
			if number > 50 {
				results <- result{0, fmt.Errorf("число %d больше 50", number)}
				return
			}

			results <- result{number * number, nil}
		}(n)
	}

	go func() {
		wg.Wait() // Ожидание счетчика
		close(results)
	}()

	for result := range results {
		if result.err != nil {
			go func() {
				// ( !!! ) Чтобы избежать блокировка оставшихся горутин
				// другой вариант: сделать буферизированный канал нужной емкости,
				// но тут у нас неизвестно, сколько значений в канале
				for range results {
				}
			}()
			fmt.Println(result.err)
			return
		}

		fmt.Println(result.value)
	}
}
