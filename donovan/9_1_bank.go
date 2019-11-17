package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

var stop = make(chan struct{})

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case <-stop:
			return
		}
	}
}

func main() {
	go teller()

	// Теперь это потокобезопасно
	go func() {
		Deposit(100)
		fmt.Println("balance = ", Balance())
	}()

	go Deposit(200)

	time.Sleep(time.Second * 2)
	fmt.Println("balance = ", Balance())
	close(stop) // останавливаем teller
}
