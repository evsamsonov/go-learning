package main

import "fmt"

func main() {
	var num int
	go func() {
		num++
	}()
	if num == 0 {
		fmt.Printf("the value is %v\n", num)
	}
}
