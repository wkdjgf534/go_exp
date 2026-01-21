package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var lock sync.Mutex

	for range 10000 {
		go func() {
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	lock.Lock()
	fmt.Println("Expected counter:", 1000)
	fmt.Println("Actual counter:", counter)
	lock.Unlock()

	// it will be solved in the next lesson
}
