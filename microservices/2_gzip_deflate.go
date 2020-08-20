package main

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	port := 8081

	http.Handle("/helloworld", NewGzipHandler(http.HandlerFunc(helloWorldHandler)))

	log.Printf("Server start on port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(rw http.ResponseWriter, _ *http.Request) {
	response := helloWorldResponse{
		Message: "Hello, world!",
	}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

type GzipHandler struct {
	next http.Handler
}

func NewGzipHandler(next http.Handler) http.Handler {
	return &GzipHandler{
		next: next,
	}
}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
		return
	}
	if strings.Contains(encodings, "deflate") {
		panic("deflate not implemented")
	}

	h.next.ServeHTTP(w, r)
}

func (h *GzipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gw := gzip.NewWriter(w)
	defer gw.Close()

	gzipResponseWriter := GzipResponseWriter{
		ResponseWriter: w,
		gw:             gw,
	}

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(gzipResponseWriter, r)
}

type GzipResponseWriter struct {
	http.ResponseWriter

	gw *gzip.Writer
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}
