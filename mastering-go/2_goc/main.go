package main

import "C"
import "fmt"

// go build -o c-shared.o -buildmode=c-shared main.go

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

func main() {

}
