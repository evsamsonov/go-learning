package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-or(append(channels[2:], orDone)...):
			}
		}()
		return orDone
	}

	// Сигнатура функции, которая мониторится
	// и при зависании рестартуется
	type startGoroutineFn func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (heartbeat <-chan interface{})

	// "Менеджер" функции
	newSteward := func(
		timeout time.Duration,
		startGoroutine startGoroutineFn,
	) startGoroutineFn {
		return func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}
				startWard := func() { // Функция для перезапуска горутины
					wardDone = make(chan interface{})
					wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2)
				}
				startWard()

				pulse := time.Tick(pulseInterval)
			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)

					for {
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat:
							continue monitorLoop
						case <-timeoutSignal:
							log.Println("steward: ward unhealthy - restarting")
							close(wardDone) // Останавливаем зависшую
							startWard()     // Запускаем снова
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()

			return heartbeat
		}
	}

	// Выше все то же, что и было 5_healing_unhealthy_goroutines.go

	orDone := func(done <-chan interface{}, c <-chan interface{}) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer close(result)
			for {
				select {
				case <-done:
					return
				case val, ok := <-c:
					if !ok {
						return
					}
					select {
					case <-done:
					case result <- val:
					}
				}
			}
		}()
		return result
	}

	bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				var stream <-chan interface{}
				select {
				case <-done:
					return
				case maybeStream, ok := <-chanStream:
					if !ok {
						return
					}
					stream = maybeStream
				}
				for val := range orDone(done, stream) {
					select {
					case <-done:
					case valStream <- val:
					}
				}
			}
		}()
		return valStream
	}

	// Позволяет выбрать ограниченное кол-во значений из генератора
	take := func(done <-chan interface{}, valuesChan <-chan interface{}, num int) <-chan interface{} {
		result := make(chan interface{})
		go func() {
			defer close(result)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case result <- <-valuesChan:
				}
			}
		}()
		return result
	}

	doWorkFn := func(done <-chan interface{}, intList ...int) (startGoroutineFn, <-chan interface{}) {
		intChanStream := make(chan (<-chan interface{}))
		intStream := bridge(done, intChanStream)
		doWork := func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
			intStream := make(chan interface{})
			heartbeat := make(chan interface{})
			go func() {
				defer close(intStream)

				select {
				case intChanStream <- intStream:
				case <-done:
					return
				}

				pulse := time.Tick(pulseInterval)

				for {
				valueLoop:
					for _, intVal := range intList {
						if intVal < 0 {
							log.Printf("negative value: %v", intVal)
							return
						}

						for {
							select {
							case <-pulse:
								select {
								case heartbeat <- struct{}{}:
								default:
								}
							case intStream <- intVal:
								continue valueLoop
							case <-done:
								return
							}
						}
					}
				}
			}()

			return heartbeat
		}
		return doWork, intStream
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	done := make(chan interface{})
	defer close(done)

	doWork, intStream := doWorkFn(done, 1, 2, -1, 3, 4, 5)
	doWorkWithSteward := newSteward(1*time.Millisecond, doWork)
	doWorkWithSteward(done, 1*time.Hour)

	for intVal := range take(done, intStream, 6) {
		fmt.Printf("Received: %d\n", intVal)
	}
}
