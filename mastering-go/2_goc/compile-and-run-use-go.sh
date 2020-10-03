#!/bin/bash

go build -o c-shared.o -buildmode=c-shared main.go
gcc -o use-go use-go.c ./c-shared.o
./use-go
