package main

import (
	"github.com/pkg/profile"
)

func main() {
	// Сгенерирует pprof файлы, по всем профилям, указанным в Start
	//
	// Просматривать информацию можно с помощью go tool pprof
	// go tool pprof ./main /var/folders/45/mtbhhpp14c71y8vv35mrycrm0000gn/T/profile871995325/cpu.pprof
	// Для генерации графического представления нужен graphviz
	// brew install graphviz
	defer profile.Start().Stop()

	doSomething()
}

func doSomething() {
	// ...
}
