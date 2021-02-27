package main

import (
	"log"
	"os"
	"testing"
)

//  go test -run=. -bench=. -benchmem

func benchWriteRandNum(b *testing.B, buffSize int) {
	fileName := "/tmp/random"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	WriteRandNum(file, 1000000, buffSize)
	b.StopTimer()

	err = os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkWriteRandNum16(b *testing.B) {
	benchWriteRandNum(b, 16)
}

func BenchmarkWriteRandNum64(b *testing.B) {
	benchWriteRandNum(b, 64)
}

func BenchmarkWriteRandNum256(b *testing.B) {
	benchWriteRandNum(b, 256)
}

func BenchmarkWriteRandNum1024(b *testing.B) {
	benchWriteRandNum(b, 1024)
}

func BenchmarkWriteRandNum2048(b *testing.B) {
	benchWriteRandNum(b, 2048)
}
