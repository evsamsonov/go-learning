package main

import (
	"log"
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
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("Steward: ward unhealthy - restarting")
						close(wardDone) // Останавливаем зависшую
						startWard()     // Запускаем снова
						continue monitorLoop
					case <-done:
						return
					}
				}
			}()

			return heartbeat
		}
	}
}
