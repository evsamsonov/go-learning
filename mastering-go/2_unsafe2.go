package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := []int{1, 2, 3, 4}
	ptr := &arr[0]
	fmt.Println(*ptr)

	memAddress := uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	//fmt.Println(*(*int)(unsafe.Pointer(memAddress)))

	for i := 0; i < len(arr)-1; i++ {
		ptr = (*int)(unsafe.Pointer(memAddress))
		fmt.Println(*ptr)
		memAddress = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	}

	// Можем обратиться за пределы массива
	memAddress = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	ptr = (*int)(unsafe.Pointer(memAddress))
	fmt.Println(*ptr)
}
