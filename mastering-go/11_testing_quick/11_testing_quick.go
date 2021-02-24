package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	type Point3D struct {
		X, Y, Z int
		S       float64
	}

	ran := rand.New(rand.NewSource(time.Now().Unix()))
	x, _ := quick.Value(reflect.TypeOf(Point3D{}), ran)
	fmt.Printf("%+v\n", x)
}

func Add(a, b int) int {
	return a + b
}
