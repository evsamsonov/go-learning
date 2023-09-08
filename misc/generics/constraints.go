package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Max(1, 2))
	fmt.Println(Max(MyInt(100), 23))
	fmt.Println(MaxUseConstraintsLib(1, 2))
	fmt.Println(MaxUseConstraintsLib("abc", "zzz"))
}

type MyInt int64

type Ordered interface {
	~int64 | // ~ означает, что все типы, которые сводятся к int64. Например MyInt
		int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64
}

func Max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxUseConstraintsLib[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
