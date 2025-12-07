package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type DBConnection struct {
	index int32
}

func main() {

	var count int32

	connectionPool := &sync.Pool{
		New: func() any {
			atomic.AddInt32(&count, 1)
			return &DBConnection{
				index: count,
			}
		},
	}

	for range 10 {
		go func() {
			con := connectionPool.Get().(*DBConnection)
			fmt.Printf("con: %v\n", con)
			time.Sleep(time.Second)
			connectionPool.Put(con)
		}()

		time.Sleep(time.Second / 2)
	}

	fmt.Scanln()
}
