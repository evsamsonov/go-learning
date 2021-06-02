package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

//  go test 3_pool_3_warm_test.go -benchtime=10s -bench=.
func init() {
	daemonStarted := startNetworkDaemonWarm()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequestWarm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

func connectToServiceWarm() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}

func serviceConnCacheWarm() *sync.Pool {
	p := &sync.Pool{
		New: connectToServiceWarm,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemonWarm() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := serviceConnCacheWarm()

		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()

		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			cachedConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(cachedConn)
			conn.Close()
		}
	}()

	return &wg
}
