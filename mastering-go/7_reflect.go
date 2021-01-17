package main

import (
	"fmt"
	"reflect"
)

type A struct {
	a int64
	b string
	c float64
}

func main() {
	x := 100
	xr := reflect.ValueOf(&x).Elem()
	fmt.Printf("Type: %s\n", xr.Type())
	fmt.Printf("Value: %d\n", xr.Int())

	a := A{}
	ar := reflect.ValueOf(&a).Elem()
	fmt.Printf("Type: %s\n", ar.Type())
	fmt.Printf("NumField: %d\n", ar.NumField())
	for i := 0; i < ar.NumField(); i++ {
		fmt.Printf("Field %d: %s\n", i, ar.Field(i).Type())
	}
}
