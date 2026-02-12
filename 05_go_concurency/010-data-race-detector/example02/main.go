package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter = 0
	mutex   sync.Mutex
	cond    = sync.NewCond(&mutex)
)

func main() {
	go producer()
	go consumer()

	time.Sleep(5 * time.Second)
}

func producer() {
	for {
		mutex.Lock()
		if counter > 0 {
			cond.Wait()
		}
		time.Sleep(1 * time.Second)
		counter++
		fmt.Printf("Incrementing counter: %v\n", counter)
		mutex.Unlock()
		cond.Signal()
	}
}

func consumer() {
	for {
		mutex.Lock()
		if counter == 0 {
			cond.Wait()
		}
		time.Sleep(1 * time.Second)
		counter--
		fmt.Printf("Decrementing counter: %v\n", counter)
		mutex.Unlock()
		cond.Signal()
	}

}
