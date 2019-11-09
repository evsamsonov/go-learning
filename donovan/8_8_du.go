package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fmt.Println(time.Now())

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var fileCount, fileBytes int64
	for size := range fileSizes {
		fileBytes += size
		fileCount++
	}

	fmt.Printf("%d files %0.2f GB\n", fileCount, float64(fileBytes)/1e9)
	fmt.Println(time.Now())
}

func walkDir(dirName string, fileSizes chan<- int64) {
	for _, entry := range fetchDirs(dirName) {
		if entry.IsDir() {
			subDir := filepath.Join(dirName, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func fetchDirs(dirName string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du %v\n", err)
		return nil
	}

	return entries
}
