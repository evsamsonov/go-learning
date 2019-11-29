package main

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	r := httptest.NewRequest("GET", "http://127.0.0.1:80/hello", nil)
	w := httptest.NewRecorder()
	handleHelloWorld(w, r)
	assert.Equal(t, "Hello, world!", string(w.Body.Bytes()))
}