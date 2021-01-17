package main

import (
	"fmt"
	"os"
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

	var f *os.File
	printMethods(f)
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	for j := 0; j < r.NumMethod(); j++ {
		fmt.Printf("%s -> %s\n", t.Method(j).Name, r.Method(j).Type())
	}
}
