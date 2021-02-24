package main

import (
	"log"
	"os"
	"runtime/trace"
)

// Show trace
// go tool trace /tmp/traceFile.out

func main() {
	f, err := os.Create("/tmp/traceFile.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// some code...
}
