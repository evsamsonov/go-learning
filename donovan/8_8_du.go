package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// -v
var verbose = flag.Bool("v", false, "Вывод промежуточных результатов")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fmt.Println(time.Now())

	fileSizes := make(chan int64)

	var waitGroup sync.WaitGroup
	for _, root := range roots {
		waitGroup.Add(1)
		go walkDir(root, &waitGroup, fileSizes)
	}

	go func() {
		waitGroup.Wait()
		close(fileSizes)
	}()

	// Периодический вывод результата работы
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	fmt.Println("test")

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

func walkDir(dirName string, waitGroup *sync.WaitGroup, fileSizes chan<- int64) {
	defer waitGroup.Done()
	for _, entry := range fetchDirs(dirName) {
		if entry.IsDir() {
			subDir := filepath.Join(dirName, entry.Name())
			waitGroup.Add(1)
			walkDir(subDir, waitGroup, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// Ограничивающий семафор для ограничений кол-ва открытых файлов
var fileSemaphore = make(chan struct{}, 20)

func fetchDirs(dirName string) []os.FileInfo {
	fileSemaphore <- struct{}{}
	defer func() { <-fileSemaphore }()

	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}
