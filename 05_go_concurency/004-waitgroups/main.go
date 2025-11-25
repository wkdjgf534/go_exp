package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	iterations := 1000000
	wg.Go(func() {
		for range iterations {
			wg.Go(func() {
				mu.Lock()
				count++
				mu.Unlock()
			})
		}
		// heavy algoritthm
		time.Sleep(10 * time.Second)
		fmt.Println("main goroutine done")
	})

	otherIterations := 100
	wg.Go(func() {
		for range otherIterations {
			wg.Go(func() {
				mu.Lock()
				count++
				mu.Unlock()
			})
		}
		// another heavy algoritthm
		time.Sleep(10 * time.Second)
		fmt.Println("main goroutine done 2")
	})

	otherIterations2 := 100
	wg.Go(func() {
		for range otherIterations2 {
			wg.Go(func() {
				mu.Lock()
				count++
				mu.Unlock()
			})
		}
		// another heavy algoritthm
		time.Sleep(10 * time.Second)
		fmt.Println("main goroutine done 3")
	})

	wg.Wait()

	fmt.Println(count)

}
