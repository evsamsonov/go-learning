package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// По соглашению переменная, которую защищают
	// мьютексом опрделеяется сразу же, после самого мьютекса
	muB     sync.Mutex
	balance int
)

func DepositM(amount int) {
	muB.Lock()
	balance += amount
	muB.Unlock()
}

func BalanceM() int {
	muB.Lock()
	// Так лучше - выполнится даже при аварийном завершении - это имеет значение, если потом будет recover
	// но не везде возможно - блокировка до конца функции
	defer muB.Unlock()
	return balance
}

func main() {
	go func() {
		DepositM(100)
		fmt.Println("balance = ", BalanceM())
	}()

	go DepositM(200)

	time.Sleep(time.Second * 2)
	fmt.Println("balance = ", BalanceM())
}
