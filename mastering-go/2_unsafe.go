package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var val int64 = 5
	var p1 = &val
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1 = ", *p1)
	fmt.Println("*p2 = ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(val)
	fmt.Println("*p2 = ", *p2)
	*p1 = 54341234
	fmt.Println(val)
	fmt.Println("*p2 = ", *p2)
}
