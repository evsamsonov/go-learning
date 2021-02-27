package main

import (
	"io"
	"log"
	"math/rand"
)

func WriteRandNum(dst io.Writer, size, bufSize int) {
	buf := make([]byte, bufSize)
	bs := bufSize
	for {
		fillBufRandNum(buf[:bs])
		_, err := dst.Write(buf[:bs])
		if err != nil {
			log.Fatal(err)
		}
		size = size - bufSize
		if size < 0 {
			break
		}
		if size < bufSize {
			bs = size
		}
	}
}

func fillBufRandNum(buf []byte) {
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(randNum(1, 1000))
	}
}

func randNum(min, max int) int {
	return rand.Intn(max-min) + min
}
