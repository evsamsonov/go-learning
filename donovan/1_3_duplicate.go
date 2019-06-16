package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	println("Для завершения ввода используйте Ctrl+D")

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", text, count)
		}
	}
}
