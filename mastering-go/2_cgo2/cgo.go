package main

// #cgo CFLAGS: -I${SRCDIR}/call_c_lib
// #cgo LDFLAGS: ${SRCDIR}/call_c.a
// #include <stdlib.h>
// #include <call_c.h>
import "C"
import "unsafe"

func main() {
	C.cHello()

	message := C.CString("Hello, world!")
	defer C.free(unsafe.Pointer(message))

	C.printMessage(message)
}
