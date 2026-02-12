package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var data int

	go func() {
		mu.Lock()
		defer mu.Unlock()
		data = 42
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		fmt.Println(data)
	}()
	fmt.Scanln()
}

// Atomic
//func main() {
//	var data atomic.Int32
//	go func() {
//		data.Add(42)
//	}()
//	go func() {
//		fmt.Println(data.Load())
//	}()
//
//	fmt.Scanln()
//}
