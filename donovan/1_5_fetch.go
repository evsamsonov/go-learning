package main

import (
	"fmt"
	"strings"

	//"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		if false == strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		response, error := http.Get(url)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v \n", error)
			os.Exit(1)
		}

		// Вариант 1 с записью данных в переменную
		result, error := ioutil.ReadAll(response.Body)
		if error != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, error)
			os.Exit(1)
		}
		fmt.Printf("%d %s\n", response.StatusCode, result)

		// Вариант 2 сразу вывести в поток вывода без выделения памяти
		//_, error = io.Copy(os.Stdout, response.Body)
		//if error != nil {
		//	fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, error)
		//	os.Exit(1)
		//}

		response.Body.Close()
	}
}
