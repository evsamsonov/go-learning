package main

import (
	"errors"
	"fmt"
	"github.com/eapache/go-resiliency/breaker"
	"time"
)

func main() {
	mainBreaker := breaker.New(3, 1, 5*time.Second)

	for {
		result := mainBreaker.Run(func() error {
			// Call some service
			time.Sleep(2 * time.Second)
			return errors.New("timeout")
		})

		switch result {
		case nil:
			fmt.Println("Success!")
		case breaker.ErrBreakerOpen:
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
