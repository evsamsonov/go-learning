package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

type helloRequest struct {
	Name string `json:"name"`
}

func main() {
	helloCalledMetric := promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_hello_called",
	})
	helloSuccessMetric := promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_hello_success",
	})
	helloFailedMetric := promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_hello_failed",
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloCalledMetric.Inc()

		var hello helloRequest
		err := json.NewDecoder(r.Body).Decode(&hello)
		if err != nil {
			helloFailedMetric.Inc()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, %s!\n", hello.Name)
		helloSuccessMetric.Inc()
	})
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8081", nil))
}
