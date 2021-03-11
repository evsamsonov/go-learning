package main

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	pathToCert := "mastering-go/13_https"
	cert, err := ioutil.ReadFile(pathToCert + "/client.crt")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	pair, err := tls.LoadX509KeyPair(pathToCert+"/client.crt", pathToCert+"/client.key")
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{pair},
				RootCAs:            certPool,
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
