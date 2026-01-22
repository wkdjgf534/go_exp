package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	counter int
)

func highPriority() {
	mutex.Lock()
	time.Sleep(100 * time.Microsecond)
	defer mutex.Unlock()
	counter += 10
}

func lowPriority() {
	mutex.Lock()
	defer mutex.Unlock()
	counter += 1
}

func main() {
	for range 1000 {
		go highPriority() // Continuously launch high priority routines
		go lowPriority()  // Start low priority routine first
	}

	// Wait for all routines (although low priority might starve)
	time.Sleep(1 * time.Second)
	fmt.Println("Final counter value:", counter)
}
