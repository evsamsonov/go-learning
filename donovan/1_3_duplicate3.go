package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// в data байтовый срез
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		// переводим байтовый срез в строку, разбиваем по переносам строк
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%s: %d\n", text, count)
		}
	}
}
