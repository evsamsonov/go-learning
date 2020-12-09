package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	r1, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("Pid: ", r1)

	r1, _, _ = syscall.Syscall(24, 0, 0, 0)
	fmt.Println("User ID: ", r1)

	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	c, err := syscall.Write(1 /* дескриптор файла */, message)
	fmt.Println(err)
	fmt.Println("Count: ", c)

	err = syscall.Exec("/bin/ls", []string{"-l"}, os.Environ())
	fmt.Println(err)
}
