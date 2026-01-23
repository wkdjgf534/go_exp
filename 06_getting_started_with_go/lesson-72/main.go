package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.RWMutex
	count int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) Get() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	counter := &Counter{count: 0}
	wg := &sync.WaitGroup{}

	wg.Add(10)
	for range 10 {
		go func() {
			for range 100 {
				counter.Inc()
			}
			wg.Done()
		}()
	}

	wg.Add(10)
	for range 100 {
		go func() {
			fmt.Println("Count: ", counter.Get())
			time.Sleep(2 * time.Millisecond)
			wg.Done()
		}()
	}

	wg.Wait()
}
