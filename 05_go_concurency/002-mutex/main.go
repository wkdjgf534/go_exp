package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int
	var mu sync.Mutex

	for range 1000 {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(count)
			count++
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println(count)
}
