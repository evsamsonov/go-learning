package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var result string
	for _, value := range os.Args[1:] {
		result += value + " "
	}

	fmt.Println(result)

	// Еще вариант
	fmt.Println(strings.Join(os.Args[1:], " "))
}
