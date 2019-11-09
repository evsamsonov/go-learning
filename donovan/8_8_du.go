package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// -v
var verbose = flag.Bool("v", true, "Вывод промежуточных результатов")

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

	fmt.Println(*verbose)

	// Периодический вывод результата работы
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var fileCount, fileBytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // указывает, что нужно выйти из цикла, помеченного меткой loop
			}

			fileBytes += size
			fileCount++
		case <-tick:
			printResult(fileCount, fileBytes)
		}
	}

	printResult(fileCount, fileBytes)

	fmt.Println(time.Now())
}

func printResult(fileCount int64, fileBytes int64) (int, error) {
	return fmt.Printf("du: %d files %0.2f GB\n", fileCount, float64(fileBytes)/1e9)
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
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}
