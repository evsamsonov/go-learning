package microservices

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type Response struct {
	Message string
}

func BenchmarkMarshal(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello!"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, data)
	}
}

func BenchmarkEncoding(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello!"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(response)
	}
}

func BenchmarkEncodingReference(b *testing.B) {
	b.ResetTimer()

	var writer = ioutil.Discard
	response := Response{Message: "Hello!"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		encoder.Encode(&response)
	}
}
