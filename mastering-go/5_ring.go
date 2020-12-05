package main

import (
	"container/ring"
	"fmt"
)

func main() {
	myRing := ring.New(10)
	fmt.Println(myRing.Len())

	printRing(myRing)

	myRing.Value = 1
	printRing(myRing)

	for i := 0; i < myRing.Len(); i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	printRing(myRing)

}

func printRing(r *ring.Ring) {
	i := 0
	r.Do(func(v interface{}) {
		i++
		if v == nil {
			fmt.Printf("%d: empty\n", i)
			return
		}
		fmt.Printf("%d: %d\n", i, v)
	})

	fmt.Println()
}
