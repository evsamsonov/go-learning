package main

import (
	"flag"
	"fmt"
	"strings"
)

// Usage:
//  -n    пропуск символа новой строки
//  -s string
//        разделитель (default " ")

var n = flag.Bool("n", false, "пропуск символа новой строки")
var sep = flag.String("s", " ", "разделитель")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
