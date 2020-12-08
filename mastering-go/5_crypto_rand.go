package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateBytes(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generatePass(size int) (string, error) {
	b, err := generateBytes(size)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func main() {
	fmt.Println(generatePass(10))
}
