package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Начало отсчета. Для отмены нажмите any key")
	tick := time.Tick(time.Second)
	for i := 10; i > 0; i-- {
		fmt.Printf("%d\r", i)
		select {
		case <-tick:
		case <-abort:
			fmt.Printf("Запуск отменен!")
			return
		}
	}

	fmt.Println("Запуск!")
}
