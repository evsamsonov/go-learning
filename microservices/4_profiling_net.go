package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// По адресу localhost:6060/debug/pprof/ будет доступна информация профилирования
	// Сохранить profile в файл
	// go tool pprof http://localhost:6060/debug/pprof/profile
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
