package main

import (
	"net/http"
)

func main() {
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8090", nil)
}
