package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestContextWithTimeout(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutCtx, cancelFunc := context.WithTimeout(req.Context(), time.Millisecond*1)
	defer cancelFunc()

	req = req.WithContext(timeoutCtx)

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}
