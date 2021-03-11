package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	pathToCert := "mastering-go/13_https"
	cert, err := ioutil.ReadFile(pathToCert + "/client.crt")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "It's an example HTTPS server!")
	})

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: cfg,
	}
	err = server.ListenAndServeTLS(pathToCert+"/server.crt", pathToCert+"/server.key")
	if err != nil {
		log.Fatal(err)
	}
}
