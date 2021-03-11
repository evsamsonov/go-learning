package main

import (
	"fmt"
	"log"
	"net/http"
)

// Генерация TLS сертификата сервера
// openssl genrsa -out server.key 2048
// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

func main() {
	pathToCert := "mastering-go/13_https"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "It's an example HTTPS server!")
	})

	err := http.ListenAndServeTLS(":8080", pathToCert+"/server.crt", pathToCert+"/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
