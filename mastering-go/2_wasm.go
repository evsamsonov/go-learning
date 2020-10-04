package main

import "net/http"

/*
	Run from mastering-go
 	$ pwd
	../go-learning/microservices
	$ go run 2_wasm.go

	Prepare WASM
 	$ pwd
	../go-learning
	$ GOOS=js GOARCH=wasm go build -o main.wasm hello-world.go
	$ mv main.wasm mastering-go/2_wasm
	$ cd mastering-go/2_wasm
	$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
*/

func main() {
	http.Handle("/", http.FileServer(http.Dir("2_wasm")))
	http.ListenAndServe(":8090", nil)
}
