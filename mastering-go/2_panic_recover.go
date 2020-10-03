package main

import "fmt"

func main() {
	calmFunc()
}

func calmFunc() {
	defer func() {
		if c := recover(); c != nil {
			fmt.Printf("Recover inside calmFunc: %v", c)
		}
	}()

	fmt.Println("before calling panickingFunc")
	panickingFunc()
	fmt.Println("after calling panickingFunc")
}

func panickingFunc() {
	fmt.Println("before panic")
	panic("panic!")
	fmt.Println("after panic")
}
