package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"log"
)

func main() {
	fileName := "6_scanner.go"
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	fileSet := token.NewFileSet()
	files := fileSet.AddFile(fileName, fileSet.Base(), len(fileContent))

	var myScanner scanner.Scanner
	myScanner.Init(files, fileContent, nil, scanner.ScanComments)

	for {
		pos, tok, lit := myScanner.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fileSet.Position(pos), tok, lit)
	}
}
