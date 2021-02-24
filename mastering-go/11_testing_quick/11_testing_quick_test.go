package main

import (
	"testing"
	"testing/quick"
)

func TestAdd(t *testing.T) {
	condition := func(a, b int) bool {
		return Add(a, b) == (a + b)
	}

	err := quick.Check(condition, &quick.Config{
		MaxCount:      1000000,
		MaxCountScale: 0,
		Rand:          nil,
		Values:        nil,
	})
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
