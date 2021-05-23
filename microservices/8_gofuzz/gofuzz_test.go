package main_test

import (
	"github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

type StructForTest struct {
	A string
	B *string
	C int
	D struct {
		E float64
	}
}

func MyFunc(o *StructForTest) error {
	// do something
	return nil
}

func TestMyFunc(t *testing.T) {
	for i := 0; i < 50000; i++ {
		o := StructForTest{}

		f := fuzz.New().NilChance(0.5)
		f.Fuzz(&o)
		err := MyFunc(&o)
		assert.Nil(t, err)
	}
}
