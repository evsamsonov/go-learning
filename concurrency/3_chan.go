package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		resultChan := make(chan int, 5) // Instance channel
		go func() {
			defer close(resultChan) // Close channel
			for i := 0; i <= 5; i++ {
				resultChan <- i // Write to channel
			}
		}()
		return resultChan // Pass chan to user
	}

	resultChan := chanOwner()
	for res := range resultChan { // Read channel and prevent block
		fmt.Printf("Received: %d\n", res)
	}
	fmt.Println("Done receiving!")
}
