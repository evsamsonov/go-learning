package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// curl -X POST -d '{"name": "Evgeny"}' http://127.0.0.1:8081/hello

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

	// see https://gist.github.com/diolavr/ef6d63288a4244b8f745958041fd3f73
	requestDurationMetric := promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "app_request_duration",
		Help: "Продолжительность обработки запроса",
	})
	requestLatencyMetric := promauto.NewSummary(prometheus.SummaryOpts{
		Name: "app_request_latency",
		Help: "Время задержки обработки запроса в секундах",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()

		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		defer func() {
			duration := time.Since(started)
			val := float64(time.Since(started) / time.Second)
			requestDurationMetric.Observe(val)
			requestLatencyMetric.Observe(val)

			log.Printf("Request duration: %s\n", duration)
		}()
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
