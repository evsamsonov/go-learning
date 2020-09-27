package main

// #include <stdio.h>
// void callC() {
//    printf("Calling C code!!!");
// }
import "C"

func main() {
	C.callC()
}
