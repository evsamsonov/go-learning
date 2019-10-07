package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Ожидание %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
