package __throttling

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func NewTestHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		<-r.Context().Done()
	})
}

func TestReturnBusy(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithCancel(context.Background())
	handler := NewLimitHandler(1, NewTestHandler())

	r := &http.Request{}
	r = r.WithContext(ctx)
	r2 := &http.Request{}
	r2 = r2.WithContext(ctx2)

	rw := &httptest.ResponseRecorder{}
	rw2 := &httptest.ResponseRecorder{}

	time.AfterFunc(10*time.Millisecond, func() {
		cancel()
		cancel2()
	})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		handler.ServeHTTP(rw, r)
	}()

	go func() {
		defer wg.Done()
		handler.ServeHTTP(rw2, r2)
	}()

	wg.Wait()

	if rw.Code == http.StatusOK && rw2.Code == http.StatusOK {
		t.Fatal("One request should be busy")
	}
}
