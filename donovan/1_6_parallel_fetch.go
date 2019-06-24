package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	netUrl "net/url" // Так можно задать алиас
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go parallelFetch(url, ch)
	}

	// Получение из канала
	// Можно писать range без присваивания
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("Total time: %.2fs\n", time.Since(start).Seconds())
}

func parallelFetch(url string, ch chan<- string) {
	start := time.Now()
	if false == strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v \n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch: read %s: %v\n", url, err)
		os.Exit(1)
	}

	// Запишем полученные данные в файл
	u, err := netUrl.Parse(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch: parse %s: %v\n", url, err)
		os.Exit(1)
	}

	fileName := strings.ReplaceAll(u.Hostname(), ".", "_")
	f, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("fetch: write %s: %v\n", fileName, err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Write(result)
	if err != nil {
		ch <- fmt.Sprintf("fetch: write %s: %v\n", url, err)
		os.Exit(1)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%s / Time: %.2fs Code: %d Lenght: %d", url, secs, resp.StatusCode, len(result))
}
