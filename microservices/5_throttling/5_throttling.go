package __throttling

import (
	"net/http"
)

type LimitHandler struct {
	connections chan struct{}
	handler     http.Handler
}

func NewLimitHandler(connCount int, next http.Handler) *LimitHandler {
	connections := make(chan struct{}, connCount)
	for i := 0; i < connCount; i++ {
		connections <- struct{}{}
	}
	return &LimitHandler{
		connections: connections,
		handler:     next,
	}
}

func (l *LimitHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	select {
	case <-l.connections:
		l.handler.ServeHTTP(rw, r)
		l.connections <- struct{}{}
	default:
		http.Error(rw, "Busy", http.StatusTooManyRequests)
	}
}
