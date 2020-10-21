package main

import (
	"bufio"
	"fmt"
	"github.com/apex/log"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: select-column column <file1> [<file2> [... <fileN>]]")
		os.Exit(1)
	}
	column, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Column value is not integer: %s", args[1])
		os.Exit(1)
	}
	if column < 0 {
		fmt.Printf("Invalid column number: %s", args[1])
		os.Exit(1)
	}

	for _, fileName := range args[2:] {
		err := processFile(fileName, column)
		if err != nil {
			fmt.Printf("process file: %s", err)
		}
	}
}

func processFile(fileName string, column int) error {
	fmt.Println("\t", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("open file: %s", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("Failed to close file %s: %s", fileName, err)
		}
	}()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("read string: %w", err)
		}

		fields := strings.Fields(line)
		if column <= len(fields) {
			fmt.Println(fields[column-1])
		}
	}
	return nil
}
