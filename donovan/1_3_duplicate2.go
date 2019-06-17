package main

import (
	"bufio"
	"fmt"
	"os"
)

// Получение данных из ввода или списка файлов
func main() {
	counts := make(map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {

		for _, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", file, err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", text, count)
		}
	}
}

// map передается по значению, но внутри map копируются ссылки, поэтому будет работать
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
