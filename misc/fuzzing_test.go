package main

import (
	"testing"
)

func Foo(i int, s string) error {
	// do something
	return nil
}

// how to run
// go test fuzzing_test.go -fuzz FuzzFoo -fuzztime 10s
func FuzzFoo(f *testing.F) {
	f.Add(5, "hello")
	f.Fuzz(func(t *testing.T, i int, s string) {
		err := Foo(i, s)
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}